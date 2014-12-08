package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"server/core/packet/scheme"
	bin "server/library/binary-parser"
	"server/modules/packets/in"
	"server/modules/packets/out"
)

type PacketS struct {
	PckAr map[uint16]scheme.PacketI
}

func (self *PacketS) NewPacket() *PacketS {
	self.PckAr = make(map[uint16]scheme.PacketI)

	// Серверные пакеты
	self.PckAr[504] = &out.ShowMobP{}
	self.PckAr[508] = &out.ActionRP{}
	self.PckAr[516] = &out.WorldP{}
	self.PckAr[517] = &out.SysMessageP{}
	self.PckAr[931] = &out.CharactersP{}
	self.PckAr[935] = &out.CreateCharRP{}
	self.PckAr[936] = &out.DeleteCharRP{}
	self.PckAr[940] = &out.CurrentDateP{}
	self.PckAr[942] = &out.ChangeSecretRP{}

	// Клиентские пакеты
	self.PckAr[1] = &in.LocalMessageP{}
	self.PckAr[6] = &in.ActionP{}
	self.PckAr[35] = &in.CheckWorldP{}
	self.PckAr[346] = &in.NewSecretP{}
	self.PckAr[347] = &in.ChangeSecretP{}
	self.PckAr[431] = &in.AuthP{}
	self.PckAr[433] = &in.EnterWorldP{}
	self.PckAr[434] = &in.ChangeCharP{}
	self.PckAr[435] = &in.CreateCharP{}
	self.PckAr[436] = &in.DeleteCharP{}
	return self
}

func (self *PacketS) Pack(packet scheme.PacketI) ([]byte, error) {
	bt, err := bin.Marshal(packet)

	lb := len(bt)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint16(lb+6))
	binary.Write(buf, binary.BigEndian, []byte{0x80, 0x00, 0x00, 0x00})
	binary.Write(buf, binary.BigEndian, bt)

	fmt.Printf("\n% x\n", buf.Bytes())

	return buf.Bytes(), err
}

func (self *PacketS) Unpack(packet scheme.PacketI, bytes []byte) error {
	err := bin.Unmarshal(bytes, packet)

	return err
}
