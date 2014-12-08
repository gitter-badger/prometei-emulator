package in

import (
	"server/core/actor"
	"server/core/packet/scheme"
	"server/modules/packets/out"
)

type EnterWorldP struct {
	Op   uint16
	Name string
}

func (self *EnterWorldP) New() scheme.PacketI {
	self.Op = 433

	return self
}

func (self *EnterWorldP) Process(player *actor.Player) scheme.PacketI {
	return new(out.WorldP).New().Process(player)
}
