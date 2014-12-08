package database

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"os"
	"server/core/config"
	"server/core/logger"
)

var (
	Engine *xorm.Engine
)

func init() {
	var err error

	Engine, _ = xorm.NewEngine("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.Get.Database.Login, config.Get.Database.Password, config.Get.Database.Name))

	err = Engine.Sync(new(Users), new(Characters))
	if err != nil {
		logger.Log.Error("Database not connected")
		os.Exit(0)
	}

	logger.Log.Info("Initializing database system")
}
