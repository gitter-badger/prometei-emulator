package scheme

import (
	"server/core/actor"
)

type PacketI interface {
	New() PacketI
	Process(player *actor.Player) PacketI
}
