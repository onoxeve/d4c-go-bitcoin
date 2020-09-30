package main

import (
	"./app/controller"
	"./app/model"
	"./bitflyer"
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

	// create api client
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	// GET: Balance
	//fmt.Println(apiClient.GetBalance())

	// GET: ticker
	ticker, _ := apiClient.GetTicker(config.Config.ProductCode)
	fmt.Println(ticker)
	fmt.Println(ticker.GetMidPrice())
	//fmt.Println(ticker.DateTime())
	//fmt.Println(ticker.TruncateDateTime(time.Hour))

	// create db connection
	fmt.Println(model.DbConnection)

	// streaming candle data(bitcoin価格)
	controller.StreamCandleData()
}
