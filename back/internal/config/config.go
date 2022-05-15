package config

import (
	"github.com/spf13/viper"
)

func LoadDBConfig(path string) (config DBConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	print(err)
	return
}

type DBConfig struct {
	Driver string `mapstructure:"DB_DRIVER"`
	Source string `mapstructure:"DB_SOURCE"`
}
