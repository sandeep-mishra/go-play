package config

import (
	"fmt"

	viper "github.com/spf13/viper"
)

type NestedURL struct {
	URL    string `mapstructure:"url"`
	Nested int    `mapstructure:"nested"`
}

type ZipCode struct {
	Zipcode string `mapstructure:"zipcode"`
}

type Element struct {
	Element []map[string]NestedURL
}

type Config struct {
	Elements []map[string]Element `mapstructure:"elements"`
	Weather  []map[string]ZipCode `mapstructure:"weather"`
}

func load() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
}
