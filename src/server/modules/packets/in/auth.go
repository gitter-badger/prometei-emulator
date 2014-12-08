package in

import (
	"fmt"
	"server/core/actor"
	"server/core/database"
	"server/core/packet/scheme"
	"server/library/crypto/cipher"
	"server/modules/packets/out"
	"strings"
	"time"
)

type AuthP struct {
	Op            uint16
	Unknown       string
	Login         string
	LenPass       uint16
	Password      [24]byte
	LenMac        uint16
	MAC           [24]byte
	UnCheat       uint16
	ClientVersion uint16

	err uint16
}

func (self *AuthP) New() scheme.PacketI {
	self.Op = 431

	return self
}

func (self *AuthP) Process(player *actor.Player) scheme.PacketI {
	user := database.Users{Name: self.Login}
	ok, _ := player.DB.Get(&user)

	chars := new(out.CharactersP)
	chars.New()

	if ok {
		hashPass := cipher.TDESEEncrypt([]byte(strings.ToUpper(user.Password[:24])), []byte(player.Date))

		if fmt.Sprintf("%x", self.Password) == fmt.Sprintf("%x", hashPass) {
			if user.State != 1 {
				player.Id = user.Id
				actor.Players.Add(player.Id, player)

				user.State = 1
				user.LastLogin = time.Now()
				player.DB.Id(user.Id).Update(&user)

				chars.Error = 0
			} else {
				chars.Error = 1104
			}
		} else {
			chars.Error = 1002
		}
	} else {
		chars.Error = 1001
	}

	return chars.Process(player)
}
