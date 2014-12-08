package in

import (
	"server/core/actor"
	"server/core/packet/scheme"
	"server/modules/packets/out"
)

type CheckWorldP struct {
	Op uint16
}

func (self *CheckWorldP) New() scheme.PacketI {
	self.Op = 35

	return self
}

func (self *CheckWorldP) Process(player *actor.Player) scheme.PacketI {
	return new(out.CheckWorldRP).New().Process(player)
}
