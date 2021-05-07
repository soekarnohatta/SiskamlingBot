package main

import (
	"SiskamlingBot/bot"
	"SiskamlingBot/bot/helper/config"
	"SiskamlingBot/bot/helper/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
	"net/http"
	"runtime"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func receiveUpdates(b *gotgbot.Bot, updater ext.Updater) {
	if config.Config.WebhookURL != "" {
		webhook := ext.WebhookOpts{
			Listen:  config.Config.WebhookListen,
			Port:    config.Config.WebhookPort,
			URLPath: config.Config.WebhookPath + b.Token,
		}

		// Delete webhook before starting the bot.
		_, err := b.DeleteWebhook(&gotgbot.DeleteWebhookOpts{DropPendingUpdates: false})
		if err != nil {
			panic("failed to delete webhook: " + err.Error())
		}

		err = updater.StartWebhook(b, webhook)
		if err != nil {
			panic("failed to start webhook: " + err.Error())
		}

		ok, err := b.SetWebhook(config.Config.WebhookURL+config.Config.WebhookPath+b.Token, &gotgbot.SetWebhookOpts{MaxConnections: 40})
		if err != nil {
			panic("failed to start webhook: " + err.Error())
		}
		if !ok {
			panic("failed to set webhook, ok is false")
		}

		log.Printf("%s has been started...\n", b.User.Username)
	} else {
		// Delete webhook before starting the bot
		_, err := b.DeleteWebhook(&gotgbot.DeleteWebhookOpts{DropPendingUpdates: false})
		if err != nil {
			panic("failed to delete webhook: " + err.Error())
		}

		err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: false})
		if err != nil {
			panic("failed to start polling: " + err.Error())
		}
		log.Printf("%s has been started...\n", b.User.Username)
	}
}

func main() {
	// Init config.
	config.NewConfig()

	// Connect to DB.
	database.NewMongo()

	// Create bot from environment value.
	b, err := gotgbot.NewBot(config.Config.BotAPIKey, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	// Add handler to dispatcher.
	bot.AddHandler(b, dispatcher)

	// Start receiving updates.
	receiveUpdates(b, updater)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
