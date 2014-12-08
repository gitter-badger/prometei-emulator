package out

import (
	"server/core/actor"
	"server/core/packet/scheme"
)

type Attribute struct {
	Id    uint16
	Value uint16
}

type Item struct {
	Id            uint16 `binary:"endian=little"`
	Num           uint16 `binary:"endian=little"`
	Durability    uint16 `binary:"endian=little"`
	MaxDurability uint16 `binary:"endian=little"`
	Energy        uint16
	MaxEnergy     uint16
	ForgeLv       uint8
	Valid         bool
	DbParam1      uint32
	DbParam2      uint32
	Attrs         [5]Attribute
	Unknown2      [119]byte
	Unknwon3      uint8
}

type Look struct {
	SynType   uint8
	Race      uint16
	BoatCheck uint8
	Items     [10]Item
	Hair      uint16 `binary:"endian=little"`
}

type Character struct {
	Flag    bool
	Name    string
	Job     string
	Level   uint16
	LenLook uint16
	Look    Look
}

type CharData struct {
	Key        string `binary:"null-off"`
	NumChar    uint8
	Character1 Character `binary:"if=NumChar=1-3"`
	Character2 Character `binary:"if=NumChar=2-3"`
	Character3 Character `binary:"if=NumChar=3"`
	Pincode    uint8
	Encryption [6]byte
	Magic      uint16 `binary:"endian=little"`
}

type CharactersP struct {
	Op       uint16
	Error    uint16
	CharData CharData `binary:"if=Error=0"`
}

func (self *CharactersP) New() scheme.PacketI {
	attr := Attribute{1, 2}
	attrs := [5]Attribute{attr, attr, attr, attr, attr}
	item := Item{0, 1, 20000, 20000, 0, 0, 0, true, 0, 0, attrs, [119]byte{}, 1}
	items := [10]Item{item, item, item, item, item, item, item, item, item, item}
	look := Look{0, 2, 0, items, 2000}
	character := Character{true, "Test", "Profession", 100, 1626, look}
	chardata := CharData{string([]byte{0x7C, 0x35, 0x09, 0x19, 0xB2, 0x50, 0xD3, 0x49}), 1, character, character, character, 1, [6]byte{}, 5170}

	self.Op = 931
	self.Error = 0
	self.CharData = chardata

	return self
}

func (self *CharactersP) Process(player *actor.Player) scheme.PacketI {
	return self
}
