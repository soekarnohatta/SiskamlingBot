package main

import (
	app "SiskamlingBot/bot/core"
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
	err := bot.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	runtime.Goexit()
}
