package main

import (
	"SiskamlingBot/bot"
	"SiskamlingBot/bot/handlers/metrics"
	"SiskamlingBot/bot/handlers/picture"
	"SiskamlingBot/bot/handlers/username"
	"SiskamlingBot/bot/helpers/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"log"
	"net/http"
	"runtime"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.LstdFlags)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func addHandler(dispatcher *ext.Dispatcher) {
	dispatcher.AddHandler(handlers.NewMessage(filters.All, metrics.ChatMetrics))
	dispatcher.AddHandler(handlers.NewMessage(filters.All, metrics.UsernameMetrics))
	dispatcher.AddHandler(handlers.NewMessage(filters.All, username.Username))
	dispatcher.AddHandler(handlers.NewMessage(filters.All, picture.Picture))

	dispatcher.AddHandlerToGroup(handlers.NewCallback(filters.Prefix("username("), username.UsernameCB), 3)
	dispatcher.AddHandlerToGroup(handlers.NewCallback(filters.Prefix("picture("), picture.PictureCB), 4)
}

func receiveUpdates(b *gotgbot.Bot, updater ext.Updater) {
	if bot.Config.WebhookURL != "" {
		webhook := ext.WebhookOpts{
			Listen:  bot.Config.WebhookListen,
			Port:    bot.Config.WebhookPort,
			URLPath: bot.Config.WebhookPath + b.Token,
		}

		// Delete webhook before starting the bot.
		_, err := b.DeleteWebhook(&gotgbot.DeleteWebhookOpts{DropPendingUpdates: true})
		if err != nil {
			panic("failed to delete webhook: " + err.Error())
		}

		err = updater.StartWebhook(b, webhook)
		if err != nil {
			panic("failed to start webhook: " + err.Error())
		}

		ok, err := b.SetWebhook(bot.Config.WebhookURL+bot.Config.WebhookPath+b.Token, &gotgbot.SetWebhookOpts{MaxConnections: 40})
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
	bot.NewConfig()

	// Connect to DB.
	database.NewMongo()

	// Create bot from environment value.
	b, err := gotgbot.NewBot(bot.Config.BotAPIKey, &gotgbot.BotOpts{
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

	// Add handlers to dispatcher.
	addHandler(dispatcher)

	// Start receiving updates.
	receiveUpdates(b, updater)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
