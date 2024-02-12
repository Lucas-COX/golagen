package internal

import (
	"github.com/spf13/viper"
)

func ReadConfigFile(p *string) (*Config, error) {
	var config *Config = nil

	if p != nil {
		viper.SetConfigFile(*p)
	} else {
		viper.SetConfigFile("./configs/golagen.yml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
