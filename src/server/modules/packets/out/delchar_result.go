package out

import (
	"server/core/actor"
	"server/core/packet/scheme"
)

type DeleteCharRP struct {
	Op    uint16
	Error uint16
}

func (self *DeleteCharRP) New() scheme.PacketI {
	self.Op = 936
	self.Error = 0

	return self
}

func (self *DeleteCharRP) Process(player *actor.Player) scheme.PacketI {
	// Не правильный секретный код
	// self.Error = 534

	return self
}
