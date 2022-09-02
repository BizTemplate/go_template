package infrastructure

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"server.port"`
	Ip   string `mapstructure:"server.ip"`
}

func LoadConfig(path string, fileName string, extension string, enableWatcher bool) (config Config, err error) {

	viper.AddConfigPath(path)      //config/application.yaml => upload bean
	viper.SetConfigName(fileName)  // Register config file name (no extension)
	viper.SetConfigType(extension) // Look for specific type

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return config, err

	}

	if enableWatcher {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
		})
	}

	err = viper.Unmarshal(&config)

	return config, err
}
