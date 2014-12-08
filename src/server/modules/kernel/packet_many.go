package kernel

import (
	"server/core/actor"
	"server/core/packet"
	"server/core/packet/scheme"
)

type PacketToMany struct {
}

func (self *PacketToMany) Check(pck *packet.PacketS) {
	select {
	case packet := <-actor.Players.OutPacket:
		for _, v := range actor.Players.Array {
			for k, v2 := range packet {
				cplayer := (*v)
				cplayer.Char.SetId(uint(k))
				b, _ := pck.Pack(v2.(scheme.PacketI).Process(&cplayer))
				v.State.Write(b)
			}
		}
	default:

	}

	return
}
