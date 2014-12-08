package out

import (
	"server/core/actor"
	"server/core/packet/scheme"
)

type ActionRP struct {
	Op      uint16
	Cid     uint32
	Counter uint32
	Group   uint8
	Angle   uint16
	Act     uint16
	Csize   uint16
	StartX  uint32 `binary:"endian=little"`
	StartY  uint32 `binary:"endian=little"`
	EndX    uint32 `binary:"endian=little"`
	EndY    uint32 `binary:"endian=little"`
}

func (self *ActionRP) New() scheme.PacketI {
	self.Op = 508
	self.Cid = 10263
	self.Counter = 1
	self.Group = 1
	self.Angle = 1
	self.Act = 1
	self.Csize = 16
	self.StartX = 229000
	self.StartY = 266300
	self.EndX = 229000
	self.EndY = 266300

	return self
}

func (self *ActionRP) Process(player *actor.Player) scheme.PacketI {
	self.Cid = uint32((*player).Char.Id)

	self.StartX = uint32((*player).Char.X)
	self.StartY = uint32((*player).Char.Y)
	self.EndX = uint32((*player).Char.X)
	self.EndY = uint32((*player).Char.Y)

	return self
}
