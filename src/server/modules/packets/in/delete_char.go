package in

import (
	"server/core/actor"
	"server/core/packet/scheme"
	"server/modules/packets/out"
)

type DeleteCharP struct {
	Op     uint16
	Name   string
	Secret string
}

func (self *DeleteCharP) New() scheme.PacketI {
	self.Op = 436

	return self
}

func (self *DeleteCharP) Process(player *actor.Player) scheme.PacketI {
	return new(out.DeleteCharRP).New().Process(player)
}
