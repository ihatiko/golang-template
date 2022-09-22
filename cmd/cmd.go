package cmd

import (
	"test/config"
	"test/internal/server"
)

func Run() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	server := server.NewServer(cfg)
	cfg.Log.SetConfiguration(cfg.Server.Name)
	server.Run()
}
