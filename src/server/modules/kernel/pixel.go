package kernel

import (
	"server/core/actor"
	"server/core/packet"
)

type Pixel struct {
	X      int
	Y      int
	Player *actor.Player
}

func (self *Pixel) PaintPixel(pck *packet.PacketS) {
	/*
		if len(actor.Players.Array) != 0 {
			grid := []Tch{}

			for _, v := range actor.Players.Array {
				if v.Id != 0 && v.Char.X != 0 && v.Char.Y != 0 {
					x := int(v.Char.X) / (10 * 2.0)
					y := int(v.Char.Y) / (10 * 2.0)

					grid = append(grid, Tch{x, y, v, false})
				}
			}

			for _, v := range grid {
				for k2, v2 := range grid {
					if len(visible) == 0 {
						visible[k2] = false
					}

					if v.Player.Id != v2.Player.Id {
						if v.X-v2.X >= -150 && v.X-v2.X <= 150 {
							if v.Y-v2.Y >= -150 && v.Y-v2.Y <= 150 {
								if visible[k2] != true {
									b, _ := pck.Pack(pck.NewPacket().PckAr[504].New().Process(v2.Player))
									v.Player.State.Write(b)

									visible[k2] = true
								}

								if len(oldgrid) >= 2 {
									if oldgrid[k2].Player.Id == v2.Player.Id {
										if oldgrid[k2].X != v2.X || oldgrid[k2].Y != v2.Y {
											c, _ := pck.Pack(pck.NewPacket().PckAr[508].New().Process(v2.Player))
											v.Player.State.Write(c)
										}
									}
								}
							} else {
								visible[k2] = false
							}
						} else {
							visible[k2] = false
						}
					}
				}
			}

			oldgrid = grid
		}
	*/
}
