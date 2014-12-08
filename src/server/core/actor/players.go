package actor

import (
	"sync"
)

var (
	Players PlayersSt
)

type PlayersSt struct {
	Array     map[int]*Player
	OutPacket chan map[int]interface{}
	lock      sync.RWMutex
}

func (self *PlayersSt) Add(value int, player *Player) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.Array[value] = player
}

func (self *PlayersSt) Delete(value int) {
	self.lock.Lock()
	defer self.lock.Unlock()

	delete(self.Array, value)
}
