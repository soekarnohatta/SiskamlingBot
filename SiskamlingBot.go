package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"SiskamlingBot/bot/core/app"
	_ "SiskamlingBot/bot/modules"
)

func main() {
	var config, err = app.NewConfig()
	if err != nil {
		panic(err)
	}

	var bot = app.NewBot(config)
	var done = make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := bot.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	<-done
	log.Println("OS Interrupt Detected, Exiting ... ")

	err = bot.Updater.Stop()
	if err != nil {
		log.Fatal(err)
	}
}
