package out

import (
	"server/core/actor"
	"server/core/packet/scheme"
)

type ShowMobP struct {
	Op             uint16
	SeeType        uint8
	CharBase       CharBase
	CharSide       CharSide
	EntEvent       EntEvent
	CharLook       CharLook
	CharPk         CharPk
	CharAppendLook CharAppendLook
	NpcType        uint8
	NpcState       uint8
	CheckNext      uint16
	CharOption     CharOption
	CharSkillState CharSkillState
}

func (self *ShowMobP) New() scheme.PacketI {
	options := [74]Option{}
	for i := 0; i < len(options); i++ {
		options[i] = Option{uint8(i), 1000}
	}

	charoption := CharOption{0, 74, options}
	states := []State{}
	charskillstate := CharSkillState{0, states}

	appendt := Append{0, true}
	appends := [4]Append{appendt, appendt, appendt, appendt}
	charappendlook := CharAppendLook{appends}
	charpk := CharPk{0}
	attr := Attribute{0, 0}
	attrs := [5]Attribute{attr, attr, attr, attr, attr}
	item := CharItem{0, 0, 0, 0, 0, 0, 0, true, 0, 0, 0, 0, attrs}
	items := [10]CharItem{item, item, item, item, item, item, item, item, item, item}
	charlook := CharLook{0, 2, 0, 2291, items}
	entevent := EntEvent{10263, 1, 0, ""}
	charside := CharSide{0}
	charbase := CharBase{4, 10263, 10263, string(charset.Encode("Свой парень")), 25346, 0, 11437, 1, string(charset.Encode("Свой парень")), "", 4, 0, 0, string(charset.Encode("Своя Гильдия")), "", "", 1, 217375, 278125, 40, 333, 0}

	self.Op = 504
	self.SeeType = 0
	self.CharBase = charbase
	self.CharSide = charside
	self.EntEvent = entevent
	self.CharLook = charlook
	self.CharPk = charpk
	self.CharAppendLook = charappendlook
	self.NpcType = 0
	self.NpcState = 0
	self.CheckNext = 0
	self.CharOption = charoption
	self.CharSkillState = charskillstate

	return self
}

func (self *ShowMobP) Process(player *actor.Player) scheme.PacketI {
	self.CharBase.CharCId = uint32((*player).Char.Id)
	self.CharBase.CharMId = uint32((*player).Char.Id)
	self.EntEvent.EnityId = uint32((*player).Char.Id)
	self.CharBase.PosX = uint32(player.Char.X)
	self.CharBase.PosY = uint32(player.Char.Y)

	return self
}
