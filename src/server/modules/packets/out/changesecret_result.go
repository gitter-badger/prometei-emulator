package out

import (
	"server/core/actor"
	"server/core/packet/scheme"
)

type ChangeSecretRP struct {
	Op    uint16
	Error uint16
}

func (self *ChangeSecretRP) New() scheme.PacketI {
	self.Op = 942
	self.Error = 0

	return self
}

func (self *ChangeSecretRP) Process(player *actor.Player) scheme.PacketI {
	// Не удалось изменить секретный код
	// self.Error = 534

	return self
}
