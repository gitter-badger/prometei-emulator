package out

import (
	"server/core/actor"
	"server/core/packet/scheme"
)

type CreateCharRP struct {
	Op    uint16
	Error uint16
}

func (self *CreateCharRP) New() scheme.PacketI {
	self.Op = 935
	self.Error = 0

	return self
}

func (self *CreateCharRP) Process(player *actor.Player) scheme.PacketI {
	return self
}
