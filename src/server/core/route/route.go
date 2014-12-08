package route

import (
	"bytes"
	"encoding/binary"
	"server/core/actor"
	"server/core/packet"
)

func Route(spack *packet.PacketS, b []byte, player *actor.Player) ([]byte, bool, int, error) {
	var (
		getb []byte
		err  error
	)

	// Пакет пинга
	if bytes.Compare(b, []byte{0x00, 0x02}) == 0 {
		return []byte{0x00, 0x02}, true, 1, nil
	}

	// Пакет выхода
	if bytes.Compare(b, []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0x01, 0x01, 0xB0}) == 0 {
		return nil, false, 0, nil
	}

	op := binary.BigEndian.Uint16(b[6:8])

	if spack.PckAr[op] != nil {

		err = spack.Unpack(spack.PckAr[op], b[6:])
		if err != nil {
			return nil, false, 1, err
		}

		getb, err = spack.Pack(spack.PckAr[op].New().Process(player))
		if err != nil {
			return nil, false, 1, err
		}

		// Ошибки на входе в аккаунт, которые приводят к закрытию соединения.
		if op == 431 {
			if binary.BigEndian.Uint16(getb[8:10]) != 0 {
				return getb, true, 0, nil
			}
		}

		return getb, true, 1, nil
	}

	return nil, false, 1, nil
}
