package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"server/core/actor"
	"server/core/config"
	"server/core/database"
	"server/core/logger"
	"server/core/packet"
	"server/core/route"
	"server/modules/kernel"
)

func main() {
	var err error

	runtime.GOMAXPROCS(runtime.NumCPU())
	logger.Log.Info("Initializing emulator MMORPG Tales of Pirates [Version 0.1.0]")

	server, err := net.Listen("tcp", config.Get.Server.IP)
	if err != nil {
		logger.Log.Error("TCP error opened")
		os.Exit(0)
	}

	players := actor.PlayersSt{}
	players.Array = make(map[int]*actor.Player)
	actor.Players = players
	actor.Players.OutPacket = make(chan map[int]interface{})

	// Запуск робота
	go Robot()

	for {
		connect, err := server.Accept()
		if err != nil {
			logger.Log.Error("Client connected error")
			os.Exit(0)
		}

		go Join(connect)
	}
}

func Robot() {
	pck := new(packet.PacketS)
	kerns := new(kernel.Kernel)

	for {
		kerns.PacketToMany.Check(pck)
	}

	/*
		kerns := new(kernel.Kernel)

		oldgrid := []Tch{}
		visible := map[int]bool{}

		for {
			kerns.PaintPixel(pck)
		}
	*/
}

func Join(connect net.Conn) {
	player := new(actor.Player)
	player.Char = actor.Character{}
	player.DB = database.Engine
	player.State = connect

	// Отправка начального пакета с датой
	pck := new(packet.PacketS)
	b, _ := pck.Pack(pck.NewPacket().PckAr[940].New().Process(player))
	connect.Write(b)

	logger.Log.Info("Client connected [IP: ", connect.RemoteAddr().String(), "]")

	connectq := make(chan bool)
	go func() {
		for {
			if <-connectq {
				player.DB.Table(new(database.Users)).Id(player.Id).Update(map[string]interface{}{"state": 0})

				actor.Players.Delete(player.Id)
				connect.Close()

				return
			}
		}
	}()

	for {
		h := make([]byte, 2048)
		bwrite, _ := connect.Read(h)

		if bwrite != 0 {
			go func() {
				bt, allow, option, err := route.Route(pck, h[:bwrite], player)
				if err != nil {
					logger.Log.Error(err)
				}

				if allow == true {
					connect.Write(bt)
				}

				if option == 0 {
					connectq <- true
				}

				//logger.Log.Info(h[:bwrite])
				fmt.Printf("\n% x\n", h[:bwrite])
			}()
		} else {
			connectq <- true
		}
	}
}
