package cmd

import (
	"github.com/ihatiko/log"
	"test/config"
)

func Run() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	log.FatalF("%s", "key")
	cfg = cfg
}
