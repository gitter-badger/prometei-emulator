package out

import (
	"server/core/actor"
	"server/core/packet/scheme"
)

type CheckWorldRP struct {
	Op      uint16
	Error   uint16
	Unknown uint8
	Max     uint16
}

func (self *CheckWorldRP) New() scheme.PacketI {
	self.Op = 554
	self.Error = 0
	self.Unknown = 16
	self.Max = 65535

	return self
}

func (self *CheckWorldRP) Process(player *actor.Player) scheme.PacketI {
	return self
}
