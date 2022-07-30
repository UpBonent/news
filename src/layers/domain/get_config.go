package domain

import (
	"github.com/UpBonent/news/src/layers/infrastructure/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

var instance *Config
var once sync.Once

func GetConfig(l *logging.Logger) *Config {
	once.Do(func() {
		l.Info("read config")

		instance = &Config{}

		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			//helpText := "news project is working"
			//_, err = cleanenv.GetDescription(instance, &helpText)
			l.Info("read config -- DONE")
		} else {
			l.Errorf("Problem with read config: %v", err)
		}

	})
	return instance
}
