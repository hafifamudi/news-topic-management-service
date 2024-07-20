package config

import (
	"github.com/jinzhu/configor"
	"sync"
)

type App struct {
	Name string `default:"news_topic_management_service" env:"APP_NAME"`
	Env  string `default:"local" env:"APP_ENV"`
}

type DBConfig struct {
	Client   string `default:"postgresql" env:"DB_CLIENT"`
	Host     string `default:"127.0.0.1" env:"DB_HOST"`
	Username string `default:"root" env:"DB_USERNAME"`
	Password string `default:"password" env:"DB_PASSWORD"`
	Port     string `default:"5432" env:"DB_PORT"`
	Database string `default:"news_topic_management_service" env:"DB_DATABASE"`
}

type Config struct {
	App App
	DB  DBConfig
}

var config *Config
var configLock = &sync.Mutex{}

func Instance() Config {
	if config == nil {
		err := Load()
		if err != nil {
			panic(err)
		}
	}
	return *config
}

func Load() error {
	tmpConfig := Config{}
	err := configor.Load(&tmpConfig)
	if err != nil {
		return err
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &tmpConfig

	return nil
}

func GetConfig() *Config {
	return config
}
