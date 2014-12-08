package binary

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type Options struct {
	Endian     binary.ByteOrder
	StringNull bool
	If         bool
}

func (self *Options) Prepare() *Options {
	self.Endian = binary.BigEndian
	self.StringNull = true
	self.If = true

	return self
}

func (self *Options) Tag(stag reflect.StructField, fd reflect.Value, FValue map[string]interface{}) error {
	tag := stag.Tag.Get("binary")

	FValue[stag.Name] = fd.Interface()

	if tag != "" {
		slice := strings.Split(tag, " ")

		for _, val := range slice {
			fval := strings.Split(val, "=")

			switch fval[0] {
			case "endian":
				if fval[1] == "little" {
					self.Endian = binary.LittleEndian
				} else {
					self.Endian = binary.BigEndian
				}
			case "null-off":
				self.StringNull = false
			case "if":
				var check int = 0

				arrnum := strings.Split(fval[2], "-")
				checkf := func(one, two interface{}, option int) {
					if option == 0 {
						if fmt.Sprintf("%v", one) == fmt.Sprintf("%v", two) {
							check++
						}
					} else {
						if fmt.Sprintf("%v", one) != fmt.Sprintf("%v", two) {
							check++
						}
					}
				}

				if len(arrnum) == 1 {
					minus := arrnum[0][:1]

					if minus == "!" {
						checkf(FValue[fval[1]], arrnum[0][1:], 1)
					} else {
						checkf(FValue[fval[1]], arrnum[0], 0)
					}
				} else {
					onenum, _ := strconv.Atoi(arrnum[0])
					twonum, _ := strconv.Atoi(arrnum[1])

					if onenum <= twonum {
						for i := onenum; i <= twonum; i++ {
							checkf(FValue[fval[1]], i, 0)
						}
					} else {
						err := errors.New("(binary-parser) In one of the conditions, the number larger than the other.")
						return err
					}
				}

				if check <= 0 {
					self.If = false
				}
			}
		}
	}

	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	b := &bytes.Buffer{}

	if err := NewEncoder(b).Encode(v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func Unmarshal(b []byte, v interface{}) error {
	return NewDecoder(bytes.NewReader(b)).Decode(v)
}

type Encoder struct {
	Option Options
	w      io.Writer
	buf    []byte
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		Option: Options{Endian: binary.BigEndian, StringNull: true},
		w:      w,
		buf:    make([]byte, 8),
	}
}

func (e *Encoder) writeVarint(v int) error {
	err := binary.Write(e.w, e.Option.Endian, uint16(v))
	return err
}

func (b *Encoder) Encode(v interface{}) (err error) {
	switch cv := v.(type) {
	case encoding.BinaryMarshaler:
		buf, err := cv.MarshalBinary()
		if err != nil {
			return err
		}
		if err = b.writeVarint(len(buf)); err != nil {
			return err
		}
		_, err = b.w.Write(buf)

	case []byte: // fast-path byte arrays
		if err = b.writeVarint(len(cv)); err != nil {
			return
		}
		_, err = b.w.Write(cv)

	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		t := rv.Type()

		switch t.Kind() {
		case reflect.Array:
			l := t.Len()
			for i := 0; i < l; i++ {
				if err = b.Encode(rv.Index(i).Addr().Interface()); err != nil {
					return
				}
			}

		case reflect.Slice:
			l := rv.Len()

			for i := 0; i < l; i++ {
				if err = b.Encode(rv.Index(i).Addr().Interface()); err != nil {
					return
				}
			}

		case reflect.Struct:
			l := rv.NumField()
			FValue := make(map[string]interface{})

			for i := 0; i < l; i++ {
				if v := rv.Field(i); v.CanSet() && t.Field(i).Name != "_" {
					if err = b.Option.Prepare().Tag(rv.Type().Field(i), v, FValue); err != nil {
						return
					}

					if b.Option.If == true {
						if err = b.Encode(v.Addr().Interface()); err != nil {
							return
						}
					}
				}
			}

		case reflect.Map:
			l := rv.Len()
			if err = b.writeVarint(l); err != nil {
				return
			}
			for _, key := range rv.MapKeys() {
				value := rv.MapIndex(key)
				if err = b.Encode(key.Interface()); err != nil {
					return err
				}
				if err = b.Encode(value.Interface()); err != nil {
					return err
				}
			}

		case reflect.String:

			var l int

			if b.Option.StringNull == true {
				l = rv.Len() + 1
				defer b.w.Write([]byte{0x00})
			} else {
				l = rv.Len()
			}

			if err = b.writeVarint(l); err != nil {
				return
			}
			_, err = b.w.Write([]byte(rv.String()))

		case reflect.Bool:
			var out byte
			if rv.Bool() {
				out = 1
			}
			err = binary.Write(b.w, b.Option.Endian, out)

		case reflect.Int:
			err = binary.Write(b.w, b.Option.Endian, int64(rv.Int()))

		case reflect.Uint:
			err = binary.Write(b.w, b.Option.Endian, int64(rv.Uint()))

		case reflect.Int8, reflect.Uint8, reflect.Int16, reflect.Uint16,
			reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128:
			err = binary.Write(b.w, b.Option.Endian, v)

		default:
			return errors.New("(binary-parser) unsupported type " + t.String())
		}
	}
	return
}

type byteReader struct {
	io.Reader
}

func (b *byteReader) ReadByte() (byte, error) {
	var buf [1]byte
	if _, err := io.ReadFull(b, buf[:]); err != nil {
		return 0, err
	}
	return buf[0], nil
}

type Decoder struct {
	Option Options
	r      *byteReader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		Option: Options{Endian: binary.BigEndian, StringNull: true},
		r:      &byteReader{r},
	}
}

func (d *Decoder) Decode(v interface{}) (err error) {
	// Check if the type implements the encoding.BinaryUnmarshaler interface, and use it if so.
	if i, ok := v.(encoding.BinaryUnmarshaler); ok {
		var l uint64
		if l, err = binary.ReadUvarint(d.r); err != nil {
			return
		}
		buf := make([]byte, l)
		_, err = d.r.Read(buf)
		return i.UnmarshalBinary(buf)
	}

	// Otherwise, use reflection.
	rv := reflect.Indirect(reflect.ValueOf(v))
	if !rv.CanAddr() {
		return errors.New("(binary-parser) can only Decode to pointer type")
	}
	t := rv.Type()

	switch t.Kind() {
	case reflect.Array:
		len := t.Len()
		for i := 0; i < int(len); i++ {
			if err = d.Decode(rv.Index(i).Addr().Interface()); err != nil {
				return
			}
		}

	case reflect.Slice:
		var l uint64
		if l, err = binary.ReadUvarint(d.r); err != nil {
			return
		}
		if t.Kind() == reflect.Slice {
			rv.Set(reflect.MakeSlice(t, int(l), int(l)))
		} else if int(l) != t.Len() {
			return fmt.Errorf("(binary-parser) encoded size %d != real size %d", l, t.Len())
		}
		for i := 0; i < int(l); i++ {
			if err = d.Decode(rv.Index(i).Addr().Interface()); err != nil {
				return
			}
		}

	case reflect.Struct:
		l := rv.NumField()
		FValue := make(map[string]interface{})

		for i := 0; i < l; i++ {
			if v := rv.Field(i); v.CanSet() && t.Field(i).Name != "_" {
				if err = d.Option.Prepare().Tag(rv.Type().Field(i), v, FValue); err != nil {
					return
				}

				if d.Option.If == true {
					if err = d.Decode(v.Addr().Interface()); err != nil {
						return
					}
				}
			}
		}

	case reflect.Map:
		var l uint64
		if l, err = binary.ReadUvarint(d.r); err != nil {
			return
		}
		kt := t.Key()
		vt := t.Elem()
		rv.Set(reflect.MakeMap(t))
		for i := 0; i < int(l); i++ {
			kv := reflect.Indirect(reflect.New(kt))
			if err = d.Decode(kv.Addr().Interface()); err != nil {
				return
			}
			vv := reflect.Indirect(reflect.New(vt))
			if err = d.Decode(vv.Addr().Interface()); err != nil {
				return
			}
			rv.SetMapIndex(kv, vv)
		}

	case reflect.String:
		var l uint16
		if err = binary.Read(d.r, d.Option.Endian, &l); err != nil {
			return
		}

		buf := make([]byte, l)
		_, err = d.r.Read(buf)
		rv.SetString(string(buf[:len(buf)-1]))

	case reflect.Bool:
		var out byte
		err = binary.Read(d.r, d.Option.Endian, &out)
		rv.SetBool(out != 0)

	case reflect.Int:
		var out int64
		err = binary.Read(d.r, d.Option.Endian, &out)
		rv.SetInt(out)

	case reflect.Uint:
		var out uint64
		err = binary.Read(d.r, d.Option.Endian, &out)
		rv.SetUint(out)

	case reflect.Int8, reflect.Uint8, reflect.Int16, reflect.Uint16,
		reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		err = binary.Read(d.r, d.Option.Endian, v)

	default:
		return errors.New("(binary-parser) unsupported type " + t.String())
	}
	return
}
