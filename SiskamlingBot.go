package main

import (
	"SiskamlingBot/bot/core/app"
	_ "SiskamlingBot/bot/modules"
	"os"
	"os/signal"
	"syscall"
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
			panic(err)
		}
	}()
	<-done
	bot.ErrorLog.Println("OS Interrupt Detected, Exiting ... ")
	err = bot.Updater.Stop()
	if err != nil {
		panic(err)
	}
}
