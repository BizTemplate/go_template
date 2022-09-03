package main

import (
	"os"
	"viper_contract/src/application"
)

func main() {
	var application = application.Application{}
	application.Start(os.Args)

}
