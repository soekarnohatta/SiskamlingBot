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
)

func main() {
	// Init config
	bot.NewConfig()

	// Connect to DB
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

	dispatcher.AddHandler(handlers.NewMessage(filters.Text, metrics.ChatMetrics))
	dispatcher.AddHandler(handlers.NewMessage(filters.Text, metrics.UsernameMetrics))

	dispatcher.AddHandlerToGroup(handlers.NewCallback(filters.Prefix("username("), username.UsernameCB), 0)
	dispatcher.AddHandlerToGroup(handlers.NewCallback(filters.Prefix("picture("), picture.PictureCB), 1)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(filters.All, username.Username), 2)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(filters.All, picture.Picture), 3)

	// Start receiving updates.
	if bot.Config.WebhookURL != "" {
		webhook := ext.WebhookOpts{
			Listen:  bot.Config.WebhookListen,
			Port:    bot.Config.WebhookPort,
			URLPath: bot.Config.WebhookPath + b.Token,
		}

		// Delete webhook before starting the bot
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
		_, err := b.DeleteWebhook(&gotgbot.DeleteWebhookOpts{DropPendingUpdates: true})
		if err != nil {
			panic("failed to delete webhook: " + err.Error())
		}

		err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
		if err != nil {
			panic("failed to start polling: " + err.Error())
		}
		log.Printf("%s has been started...\n", b.User.Username)
	}

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
