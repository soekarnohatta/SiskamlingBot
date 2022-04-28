package main

import (
	"SiskamlingBot/bot/core/app"
	"runtime"

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
