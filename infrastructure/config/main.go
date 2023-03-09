package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string `mapstructure:"TG_TOKEN"`

	DatabaseUrl string `mapstructure:"DB_URL"`

	Environment string `mapstructure:"ENVIRONMENT"`
	IsDebug     bool   `mapstructure:"DEBUG"`
}

func New(path string) (config *Config) {
	curDir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	viper.AddConfigPath(path)
	viper.SetConfigFile(curDir + "/.env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err.Error())
	}

	viper.WatchConfig()
	return config
}
