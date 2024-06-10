package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path/filepath"
	"sync"
)

const (
	LocalEnv = "local"
)

var once sync.Once

type (
	Config struct {
		Http HttpConfig
	}

	HttpConfig struct {
		HTTPAddr string `yaml:"HTTPAddr" env-required:"true"`
	}
)

// MustLoad load config. Once
func MustLoad(configDir string) *Config {
	var cfg Config

	once.Do(
		func() {
			err := cleanenv.ReadConfig(filepath.Join(configDir, Env()+".yml"), &cfg)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
		},
	)

	return &cfg
}

func Env() string {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = LocalEnv
	}

	return appEnv
}
