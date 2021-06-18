package main

import (
	"SiskamlingBot/bot/core/app"
	_ "SiskamlingBot/bot/module"
	"log"
	"runtime"
)

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	config := app.NewConfig()
	bot := app.NewBot(config)
	bot.Run()
	runtime.Goexit()
}
