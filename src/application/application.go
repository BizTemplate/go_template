package application

import (
	"log"
	"viper_contract/src/core"
	"viper_contract/src/infrastructure"
)

type Application struct {
	core.Starter
	infrastructure.Config
}

func (app *Application) Start(args []string) {
	if args[1] == "-conf" {
		var configPath = args[2]
		config, err := infrastructure.LoadFromPath(configPath)
		if err != nil {
			log.Println("Read Config Failed Failed")
			return
		}
		starter, err := core.NewStarter(config)
		if err == nil {
			starter.Start()
			return
		}
		return
	}
	log.Println("Config  not found please add cli argument -conf path")

}
