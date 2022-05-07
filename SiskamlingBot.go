package main

import (
	"SiskamlingBot/bot/core/app"
	_ "SiskamlingBot/bot/modules"
)

func main() {
	config, err := app.NewConfig()
	if err != nil {
		panic(err)
	}

	bot := app.NewBot(config)
	if err := bot.Run(); err != nil {
		panic(err)
	}
}
