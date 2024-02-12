package internal

import (
	"fmt"
	"reflect"

	"github.com/samber/lo"
	"github.com/spf13/viper"
)

func ReadConfigFile(path *string) (*Config, error) {
	var config *Config = nil
	viper.SetConfigName("configs/golagen")
	viper.SetConfigType("yaml")

	if path != nil {
		viper.AddConfigPath(*path)
	} else {
		viper.AddConfigPath(".")
	}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func checkConfigEntry(entry map[string]interface{}) error {
	var required []string = []string{"name", "route", "methods"}
	var keys []string = lo.Keys[string, interface{}](entry)

	for _, key := range required {
		_, match := lo.Find[string](keys, func(item string) bool {
			return key == item
		})
		if !match {
			return fmt.Errorf(key)
		}
	}
	return nil
}

func CheckConfigFile() error {
	var required []string = []string{"author", "project", "entries"}

	for _, key := range required {
		if !viper.InConfig(key) {
			return fmt.Errorf("missing required key \"%s\"", key)
		}
		fmt.Printf("%s: %s\n", key, reflect.TypeOf(viper.Get(key)))
	}
	entries := viper.Get("entries").([]interface{})
	for i, entry := range entries {
		err := checkConfigEntry(entry.(map[string]interface{}))
		if err != nil {
			return fmt.Errorf("missing required key \"entries[%d].%s\"", i, err.Error())
		}
	}
	return nil
}
