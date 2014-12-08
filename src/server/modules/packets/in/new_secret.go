package in

import (
	"server/core/actor"
	"server/core/packet/scheme"
)

type NewSecretP struct {
	Op     uint16
	Secret string
}

func (self *NewSecretP) New() scheme.PacketI {
	self.Op = 346

	return self
}

func (self *NewSecretP) Process(player *actor.Player) scheme.PacketI {
	return self
}
