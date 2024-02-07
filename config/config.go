package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Instance *Config

type Config struct {
	AppPort string `yaml:"AppPort"`

	Log struct {
		Path string `yaml:"Path"`
	} `yaml:"Log"`

	Baidu struct {
		APIKey    string `yaml:"apikey"`
		APISecret string `yaml:"apiSecret"`
	} `yaml:"Baidu"`

	Xunfei struct {
		AppID     string `yaml:"app_id"`
		APIKey    string `yaml:"api_key"`
		APISecret string `yaml:"api_secret"`
	} `yaml:"Xunfei"`

	Genemi struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"Genemi"`
}

func Init(filename string) *Config {
	Instance = &Config{}

	config := viper.New()
	config.SetConfigName("config") // name of config file (without extension)
	config.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath(filename) // optionally look for config in the working directory
	err := config.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := config.Unmarshal(Instance); err != nil {
		fmt.Println(err)
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
	})

	fmt.Printf("config init %v", Instance)

	return Instance
}
