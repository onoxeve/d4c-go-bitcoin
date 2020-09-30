package controller

import (
	"../../bitflyer"
	"../../config"
	"../model"
	"log"
)

// 価格チャートをストリーミング(DB保存)
func StreamCandleData() {
	var tickerChannel = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)

	for ticker := range tickerChannel {
		log.Printf("action=StreamCandleData, %v", ticker)

		for _, duration := range config.Config.Durations {
			model.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
		}
	}
}
