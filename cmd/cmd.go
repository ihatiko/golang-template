package cmd

import "test/config"

func Run() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	cfg = cfg
}
