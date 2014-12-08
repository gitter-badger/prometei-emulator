package actor

import (
	"github.com/go-xorm/xorm"
	"net"
)

type Player struct {
	Id    int
	Char  Character
	DB    *xorm.Engine
	State net.Conn
	Date  string
}
