package main

import (
	"SiskamlingBot/bot/core"
	_ "SiskamlingBot/bot/module"
	"log"
	"runtime"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	err, config := core.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	bot := core.NewBot(config)
	err = bot.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	runtime.Goexit()
}
