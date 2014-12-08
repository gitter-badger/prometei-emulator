package in

import (
	"server/core/actor"
	"server/core/packet/scheme"
	"server/modules/packets/out"
)

type ChangeCharP struct {
	Op uint16
}

func (self *ChangeCharP) New() scheme.PacketI {
	self.Op = 434

	return self
}

func (self *ChangeCharP) Process(player *actor.Player) scheme.PacketI {
	return new(out.CharactersP).New().Process(player)
}
