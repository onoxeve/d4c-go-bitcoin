package main

import (
	"./config"
	"./util"
	"fmt"
	"log"
)

func main() {
	// debug config settings
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)

	// debug logging settings
	util.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
