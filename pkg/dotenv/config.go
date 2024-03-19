package dotenv

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST               string        `mapstructure:"DB_HOST"`
	DB_PORT               string        `mapstructure:"DB_PORT"`
	DB_USER               string        `mapstructure:"DB_USER"`
	DB_PASSWORD           string        `mapstructure:"DB_PASSWORD"`
	DB_NAME               string        `mapstructure:"DB_NAME"`
	DB_DRIVER             string        `mapstructure:"DB_DRIVER"`
	ServerAddress         string        `mapstructure:"SERVER_ADDRESS"`
	JWT_SECRET            string        `mapstructure:"JWT_SECRET"`
	ACCESS_TOKEN_DURATION time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	TOKEN_SYMETRIC_KEY    string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(fmt.Sprintf("%s/.env", path))
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
