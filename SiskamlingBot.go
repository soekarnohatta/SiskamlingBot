package main

import (
	"SiskamlingBot/bot/core"
	_ "SiskamlingBot/bot/module"
	"log"
	"runtime"
)

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	config := core.NewConfig()
	bot := core.NewBot(config)
	err := bot.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	runtime.Goexit()
}
