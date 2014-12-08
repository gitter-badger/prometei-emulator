package actor

import (
	"sync"
)

type Character struct {
	Id    uint
	Hp    uint
	Mp    uint
	Level uint
	Exp   uint
	X, Y  uint
	lock  sync.RWMutex
}

func (self *Character) SetId(value uint) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.Id = value
}

func (self *Character) SetHp(value uint) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.Hp = value
}

func (self *Character) SetMp(value uint) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.Mp = value
}

func (self *Character) SetLevel(value uint) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.Level = value
}

func (self *Character) SetExp(value uint) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.Exp = value
}

func (self *Character) SetX(value uint) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.X = value
}

func (self *Character) SetY(value uint) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.Y = value
}
