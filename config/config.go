package config

import (
	"errors"
	cfg "github.com/ihatiko/viper-env"
	"github.com/spf13/viper"
	"log"
	"path"
	"time"
)

type Config struct {
	Server *Server
}

type Server struct {
	AppVersion        string
	Port              string
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

const (
	localConfigPath      = "config.yml"
	productionConfigPath = "production.yml"
	configFolder         = "config"
)

func GetConfig() (*cfg.Config, error) {
	path := path.Join(configFolder)
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

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*cfg.Config, error) {
	var c cfg.Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
