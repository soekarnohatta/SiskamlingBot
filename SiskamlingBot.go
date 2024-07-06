package main

import (
	"SiskamlingBot/bot/core/app"
	"os"
	"os/signal"
	"syscall"

	_ "SiskamlingBot/bot/modules"
)

func main() {
	config, err := app.NewConfig()
	if err != nil {
		panic(err)
	}

	bot := app.NewBot(config)
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := bot.Run(); err != nil {
			bot.ErrorLog.Fatal(err)
		}
	}()

	<-done
	bot.ErrorLog.Println("OS Interrupt Detected, Exiting ... ")

	err = bot.Updater.Stop()
	if err != nil {
		bot.ErrorLog.Fatal(err)
	}
}
