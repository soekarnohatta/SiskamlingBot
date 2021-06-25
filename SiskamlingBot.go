package main

import (
	"SiskamlingBot/bot/core/app"
	_ "SiskamlingBot/bot/modules"
	"runtime"
)

func main() {
	config := app.NewConfig()
	bot := app.NewBot(config)
	bot.Run()
	runtime.Goexit()
}
