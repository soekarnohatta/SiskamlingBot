package main

import (
	"SiskamlingBot/bot/core/app"
	_ "SiskamlingBot/bot/modules"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := app.NewConfig()
	bot := app.NewBot(config)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := bot.Run(); err != nil {
			log.Fatal(err)
		}
	}()
	<-done
	log.Print("OS Interrupt Detected, Exiting ... ")
	defer bot.SendLogMessage("Shut Down", nil)
}
