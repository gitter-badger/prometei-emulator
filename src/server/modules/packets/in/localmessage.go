package in

import (
	"server/core/actor"
	"server/core/packet/scheme"
	"server/modules/packets/out"
)

type LocalMessageP struct {
	Op      uint16
	Message string
}

func (self *LocalMessageP) New() scheme.PacketI {
	self.Op = 1

	return self
}

func (self *LocalMessageP) Process(player *actor.Player) scheme.PacketI {
	message := new(out.LocalMessageRP)
	message.New()

	message.Message = self.Message

	actor.Players.OutPacket <- map[int]interface{}{int((*player).Char.Id): message}

	return self
}
