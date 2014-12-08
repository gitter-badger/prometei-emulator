package out

import (
	"fmt"
	"server/core/actor"
	"server/core/packet/scheme"
	"time"
)

type CurrentDateP struct {
	Op   uint16
	Date string
}

func (self *CurrentDateP) New() scheme.PacketI {
	current := time.Now()

	self.Op = 940
	self.Date = fmt.Sprintf("[%d-%d %d:%d:%d:%d]", current.Month(), current.Day(), current.Hour(), current.Minute(), current.Second(), current.Nanosecond()/1000000)

	return self
}

func (self *CurrentDateP) Process(player *actor.Player) scheme.PacketI {
	player.Date = self.Date

	return self
}
