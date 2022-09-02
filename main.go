package main

import (
	"log"
	"viper_contract/src/infrastructure"
)

func main() {
	config, err := infrastructure.LoadConfig("./config", "application", "yml", true)

	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	port := config.Port

	serverAddress := config.Ip + ":" + port

	go func() {

	}()

}
