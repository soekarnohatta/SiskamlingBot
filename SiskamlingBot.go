package main

import (
	"runtime"

	"SiskamlingBot/bot/core/app"
	_ "SiskamlingBot/bot/modules"
)

func main() {
	run()
}

func run() {
	config := app.NewConfig()
	bot := app.NewBot(config)
	bot.Run()
	runtime.Goexit()
}
