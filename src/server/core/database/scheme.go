package database

import (
	"time"
)

type Users struct {
	Id        int    `xorm:"autoincr pk"`
	Name      string `xorm:"varchar(14) not null unique"`
	Password  string `xorm:"varchar(32) not null"`
	State     int
	LastLogin time.Time
}

type Characters struct {
	AccountId int    `xorm:"not null"`
	Id        int    `xorm:"autoincr pk"`
	Name      string `xorm:"varchar(14) not null unique"`
	Job       string `xorm:"varchar(14) not null"`
	Race      int    `xorm:"not null"`
}
