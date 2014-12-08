package config

import (
	toml "github.com/achun/tom-toml"
	"os"
	"server/core/logger"
)

type Config struct {
	Database struct {
		Login    string
		Password string
		Name     string
	}
	Server struct {
		IP string
	}
}

var (
	Get Config
)

func init() {
	conf, err := toml.LoadFile("resource/config/main.toml")
	if err != nil {
		logger.Log.Error("Config error loading")
		os.Exit(0)
	}

	confdb := conf.Fetch("config.database")
	confdb.Apply(&Get.Database)

	confserver := conf.Fetch("config.server")
	confserver.Apply(&Get.Server)

	logger.Log.Info("Initializing config system")
}
