package main

import "github.com/ngenator/crypto/gdax"

func main() {
	f := gdax.NewFeedWatcher()
	f.Watch("LTC-BTC", "LTC-USD", "ETH-BTC", "ETH-USD", "BTC-USD")
}
