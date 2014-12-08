package in

import (
	"server/core/actor"
	"server/core/packet/scheme"
	"server/modules/packets/out"
)

type CreateCharP struct {
	Op      uint16
	Name    string
	Map     string
	LenLook uint16
	Look    out.Look
}

func (self *CreateCharP) New() scheme.PacketI {
	self.Op = 435

	return self
}

func (self *CreateCharP) Process(player *actor.Player) scheme.PacketI {
	return new(out.CreateCharRP).New().Process(player)
}
