package config

import (
	"errors"
	"fmt"
	"github.com/ihatiko/log"
	cfg "github.com/ihatiko/viper-env"
	"github.com/spf13/viper"
	lg "log"
	"os"
	"time"
)

type Config struct {
	Server *Server
	Log    *log.Config
}

type Server struct {
	Name              string
	Version           string
	Port              string
	PprofPort         string
	Mode              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
	Debug             bool
}

func GetConfig() (*Config, error) {
	config := localConfigPath
	stand := os.Getenv(standENV)
	if stand == productionStand {
		config = productionConfigPath
	}
	path := fmt.Sprintf("%s/%s", configFolder, config)
	cfgFile, err := LoadConfig(path)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadConfig(filename string) (*cfg.Config, error) {
	cfg := cfg.New(viper.New())
	cfg.SetConfigName(filename)
	cfg.AddConfigPath(".")
	cfg.AutomaticEnv()
	if err := cfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return cfg, nil
}

func ParseConfig(v *cfg.Config) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		lg.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
