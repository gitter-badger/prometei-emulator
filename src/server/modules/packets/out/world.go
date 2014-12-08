package out

import (
	char "github.com/davidmz/go-charset"
	"math/rand"
	"server/core/actor"
	"server/core/packet/scheme"
)

var (
	charset *char.Charset = char.CP1251
)

type Shortcut struct {
	Type   uint8
	GridId uint16
}

type CharShortcut struct {
	Shortcuts [36]Shortcut
}

type Bag struct {
	GridId    uint16
	ItemId    uint16       `binary:"endian=little"`
	Num       uint16       `binary:"endian=little if=ItemId=!0"`
	Endure    uint16       `binary:"endian=little if=ItemId=!0"`
	MaxEndure uint16       `binary:"endian=little if=ItemId=!0"`
	Energy    uint16       `binary:"endian=little if=ItemId=!0"`
	MaxEnergy uint16       `binary:"endian=little if=ItemId=!0"`
	ForgeLv   uint8        `binary:"endian=little if=ItemId=!0"`
	Valid     bool         `binary:"endian=little if=ItemId=!0"`
	DbParam0  uint32       `binary:"endian=little if=ItemId=!0"`
	DbParam1  uint32       `binary:"endian=little if=ItemId=!0"`
	CheckNext uint8        `binary:"endian=little if=ItemId=!0"`
	Attrs     [5]Attribute `binary:"if=ItemId=!0 if=CheckNext=!0"`
}

type CharKitbag struct {
	Type      uint8
	KeybagNum uint16 `binary:"if=Type=0"`
	Bags      []Bag
}

type Option struct {
	AttrId uint8
	Value  uint32
}

type CharOption struct {
	Type    uint8
	Num     uint16
	Options [74]Option `binary:"if=Num=!0"`
}

type State struct {
	Id uint8
	Lv uint8
}

type CharSkillState struct {
	Num    uint8
	States []State `binary:"if=Num=!0"`
}

type Skill struct {
	Id         uint16 `binary:"endian=little"`
	State      uint8
	Lv         uint8
	UseSp      uint16 `binary:"endian=little"`
	UseEndure  uint16 `binary:"endian=little"`
	UseEnergy  uint16 `binary:"endian=little"`
	ResumeTime uint32
	Range      [4][2]byte
}

type CharSkillBag struct {
	Id     uint16 `binary:"endian=little"`
	Type   uint8
	Num    uint16  `binary:"endian=little"`
	Skills []Skill `binary:"if=Num=!0"`
}

type Append struct {
	LookId uint16 `binary:"endian=little"`
	Valid  bool   `binary:"if=LookId=!0"`
}

type CharAppendLook struct {
	Appends [4]Append
}

type CharPk struct {
	PkCtrl uint8
}

type CharItem struct {
	Id         uint16       `binary:"endian=little"`
	Num        uint16       `binary:"endian=little if=Id=!0"` //if=SynType=!1
	Endure     uint16       `binary:"endian=little if=Id=!0"`
	MaxEndure  uint16       `binary:"endian=little if=Id=!0"` //if=SynType=!1
	Energy     uint16       `binary:"endian=little if=Id=!0"`
	MaxEnergy  uint16       `binary:"endian=little if=Id=!0"` //if=SynType=!1
	ForgeLv    uint8        `binary:"endian=little if=Id=!0"` //if=SynType=!1
	Valid      bool         `binary:"if=Id=!0"`
	CheckNext1 uint8        `binary:"if=Id=!0"`                                                 //if=SynType=!1
	DbParam1   uint32       `binary:"endian=little if=Id=!0 if=CheckNext1=!0"`                  //if=SynType=!1
	DbParam2   uint32       `binary:"endian=little if=Id=!0 if=CheckNext1=!0"`                  //if=SynType=!1
	CheckNext2 uint8        `binary:"if=Id=!0"`                                                 //if=SynType=!1
	Attrs      [5]Attribute `binary:"endian=little if=Id=!0 if=CheckNext1=!0 if=CheckNext2=!0"` //if=SynType=!1
}

type CharLook struct {
	SynType   uint8
	Race      uint16
	BoatCheck uint8
	HairId    uint16
	Items     [10]CharItem
}

type EntEvent struct {
	EnityId   uint32
	EnityType uint8
	EventId   uint16
	EventName string
}

type CharSide struct {
	Id uint8
}

type CharBase struct {
	WorldId    uint32
	CharCId    uint32
	CharMId    uint32
	CharCName  string
	Unknown    uint16
	GmLv       uint8
	Handle     uint16
	CtrlType   uint8
	CharMName  string
	MottoName  string
	Icon       uint16
	Unknown2   uint16
	GuildId    uint16
	GuildName  string
	GuildMotto string
	StallName  string
	State      uint16
	PosX       uint32
	PosY       uint32
	PosRadius  uint32
	PosAngle   uint16
	TeamLId    uint32
}

type WorldP struct {
	Op             uint16
	EnterRet       uint16
	AutoLock       uint8
	KitbagLock     uint8
	EnterType      uint8
	NewChar        uint8
	MapName        string
	CanTeam        uint8
	CharBase       CharBase
	CharSide       CharSide
	EntEvent       EntEvent
	CharLook       CharLook
	CharPk         CharPk
	CharAppendLook CharAppendLook
	CharSkillBag   CharSkillBag
	CharSkillState CharSkillState
	CharOption     CharOption
	CharKitbag     CharKitbag
	MapVisible     uint16
	CharShortcut   CharShortcut
	BoatNum        uint8
	CharMainId     uint32
}

func (self *WorldP) New() scheme.PacketI {
	shortcut := Shortcut{255, 65280}
	shortcuts := [36]Shortcut{shortcut}
	charshortcut := CharShortcut{shortcuts}
	attrb := Attribute{0, 0}
	attrsb := [5]Attribute{attrb, attrb, attrb, attrb, attrb}
	bag := Bag{0, 0, 0, 0, 0, 0, 0, 0, true, 0, 0, 0, attrsb}
	bags := []Bag{bag, bag, bag, bag}
	charkitbag := CharKitbag{0, 4, bags}
	options := [74]Option{}
	for i := 0; i < len(options); i++ {
		options[i] = Option{uint8(i), 1000}
	}

	charoption := CharOption{0, 74, options}
	states := []State{}
	charskillstate := CharSkillState{0, states}
	skills := []Skill{}
	charskillbag := CharSkillBag{0, 0, 0, skills}
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
	charbase := CharBase{2, 10263, 10263, string(charset.Encode("Свой парень")), 25346, 0, 11437, 1, string(charset.Encode("Свой парень")), "", 4, 0, 0, string(charset.Encode("Своя Гильдия")), "", "", 1, 217375, 278125, 40, 333, 0}

	self.Op = 516
	self.EnterRet = 0
	self.AutoLock = 0
	self.KitbagLock = 0
	self.EnterType = 1
	self.NewChar = 0
	self.MapName = "garner"
	self.CanTeam = 1
	self.CharBase = charbase
	self.CharSide = charside
	self.EntEvent = entevent
	self.CharLook = charlook
	self.CharPk = charpk
	self.CharAppendLook = charappendlook
	self.CharSkillBag = charskillbag
	self.CharSkillState = charskillstate
	self.CharOption = charoption
	self.CharKitbag = charkitbag
	self.MapVisible = 65535
	self.CharShortcut = charshortcut
	self.BoatNum = 0
	self.CharMainId = 10263

	return self
}

func (self *WorldP) Process(player *actor.Player) scheme.PacketI {
	newid := uint(self.CharBase.CharCId) + uint(rand.Intn(100))

	self.CharBase.CharCId = uint32(newid)
	self.CharBase.CharMId = uint32(newid)
	self.EntEvent.EnityId = uint32(newid)
	self.CharMainId = uint32(newid)

	player.Char.SetId(newid)
	player.Char.SetX(uint(self.CharBase.PosX))
	player.Char.SetY(uint(self.CharBase.PosY))

	return self
}
