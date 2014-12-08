package in

import (
	"server/core/actor"
	"server/core/packet/scheme"
	"server/modules/packets/out"
)

type ChangeSecretP struct {
	Op        uint16
	SecretOld string
	SecretNew string
}

func (self *ChangeSecretP) New() scheme.PacketI {
	self.Op = 347

	return self
}

func (self *ChangeSecretP) Process(player *actor.Player) scheme.PacketI {
	return new(out.ChangeSecretRP).New().Process(player)
}
