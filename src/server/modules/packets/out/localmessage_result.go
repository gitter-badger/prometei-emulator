package out

import (
	"server/core/actor"
	"server/core/packet/scheme"
)

type LocalMessageRP struct {
	Op      uint16
	Cid     uint32
	Message string
}

func (self *LocalMessageRP) New() scheme.PacketI {
	self.Op = 501
	self.Message = string(charset.Encode("Сегодня отличная погодка"))

	return self
}

func (self *LocalMessageRP) Process(player *actor.Player) scheme.PacketI {
	self.Cid = uint32((*player).Char.Id)

	return self
}
