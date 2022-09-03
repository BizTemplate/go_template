package infrastructure

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
)

type Config struct {
	StarterConfig struct {
		IsWatchConfigFile bool   `json:"is_watch_config_file" :"is_watch_config_file"`
		StarterType       string `json:"starterType"`
	}
	ServerConfig struct {
		Port string `json:"port"`
		Ip   string `json:"ip"`
	}
}

func LoadFromPath(path string) (config Config, err error) {
	dir, file := filepath.Split(path)
	extension := filepath.Ext(path)

	return LoadConfig(filepath.Clean(dir), filepath.Clean(file), strings.ReplaceAll(extension, ".", ""))
}

func LoadConfig(path string, fileName string, extension string) (config Config, err error) {
	config, err = readConfig(path, fileName, extension)
	if err != nil {
		log.Println("Load Config Failed ", err)
		return Config{}, err
	}
	enableWatcher(config)
	return config, err
}

func enableWatcher(config Config) {
	if config.StarterConfig.IsWatchConfigFile {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
		})
		log.Println("Enable Watcher for Starter")
	}
}

func readConfig(path string, fileName string, extension string) (config Config, err error) {
	viper.AddConfigPath(path)      //config/application.yaml => upload bean
	viper.SetConfigName(fileName)  // Register config file name (no extension)
	viper.SetConfigType(extension) // Look for specific type

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return config, err

	}

	err = viper.Unmarshal(&config)
	return config, err
}
