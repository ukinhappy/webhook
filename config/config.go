package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/ukinhappy/webhook/hook"
)

func Init(path string) {
	viper.SetConfigType("toml")
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		hook.LoadAllWebHook()
	})
}
