package out

import (
	"fmt"
	"server/core/actor"
	"server/core/packet/scheme"
)

var (
	check int = 0
)

type SysMessageP struct {
	Op      uint16
	Message string
}

func (self *SysMessageP) New() scheme.PacketI {
	check++
	self.Op = 517
	self.Message = string(charset.Encode(fmt.Sprintf("%s %d", "Сегодня отличная погодка", check)))

	return self
}

func (self *SysMessageP) Process(player *actor.Player) scheme.PacketI {
	return self
}
