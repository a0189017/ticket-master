package config

import (
	"log"
	"sync"

	"github.com/jinzhu/configor"
)

var config *Config
var once sync.Once

type Server struct {
	Host string `env:"HOST" default:"0.0.0.0"`
	Port string `env:"PORT" default:"8010"`
}

type DB struct {
	Driver        string `env:"DB_DRIVER" default:"postgres"`
	User          string `env:"DB_USER" default:"admin"`
	Password      string `env:"DB_PASSWORD"`
	Host          string `env:"DB_HOST" default:"localhost"`
	Port          string `env:"DB_PORT" default:"5432"`
	Name          string `env:"DB_NAME" default:"db"`
	SlowThreshold string `env:"DB_SLOW_THRESHOLD" default:"1000" yaml:"slow_threshold"`
	Log           struct {
		Colorful string `env:"DB_LOG_COLORFUL" default:"false"`
	}
}

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
	Debug  bool   `env:"DEBUG" default:"true"`
}

func GetConfig() *Config {
	once.Do(func() {
		configFilePath := "config.yml"
		config = &Config{}
		err := configor.Load(config, configFilePath)
		if err != nil {
			log.Fatal(err)
		}

	})
	return config
}
