package in

import (
	"server/core/actor"
	"server/core/packet/scheme"
	"server/modules/packets/out"
)

type ActionP struct {
	Op      uint16
	Cid     uint32
	Counter uint32
	Group   uint8
	Csize   uint16
	StartX  uint32 `binary:"endian=little"`
	StartY  uint32 `binary:"endian=little"`
	EndX    uint32 `binary:"endian=little"`
	EndY    uint32 `binary:"endian=little"`
}

func (self *ActionP) New() scheme.PacketI {
	self.Op = 6

	return self
}

func (self *ActionP) Process(player *actor.Player) scheme.PacketI {
	action := new(out.ActionRP)
	action.New()

	player.Char.SetX(uint(self.EndX))
	player.Char.SetY(uint(self.EndY))

	return action.Process(player)
}
