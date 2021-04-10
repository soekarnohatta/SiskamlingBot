package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/soekarnohatta/Siskamling/bot"
	"github.com/soekarnohatta/Siskamling/bot/handlers/picture"
	"github.com/soekarnohatta/Siskamling/bot/handlers/username"
	"log"
	"net/http"
)

func main() {
	// Init config
	bot.NewConfig()

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

	dispatcher.AddHandlerToGroup(handlers.NewCallback(filters.Prefix("username("), username.UsernameCB), 0)
	dispatcher.AddHandlerToGroup(handlers.NewCallback(filters.Prefix("picture("), picture.PictureCB), 1)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(filters.All, username.Username), 2)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(filters.All, picture.Picture), 3)


	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
