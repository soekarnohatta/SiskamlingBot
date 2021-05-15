package core

import (
	"SiskamlingBot/bot/core/telegram"
	"log"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"go.mongodb.org/mongo-driver/mongo"
)

type TelegramBot struct {
	Bot        *gotgbot.Bot
	Dispatcher *ext.Dispatcher
	Updater    ext.Updater
	Context    *ext.Context

	Modules   map[string]Module
	Commands  map[string]telegram.Command
	Messages  map[string]telegram.Message
	Callbacks map[string]telegram.Callback

	*Config
	DB *mongo.Database
	//maxCmdSegLen int
}

func NewBot(config *Config) *TelegramBot {
	return &TelegramBot{
		Context: nil,
		Config:  config,
		Modules: make(map[string]Module),

		Commands:  make(map[string]telegram.Command),
		Messages:  make(map[string]telegram.Message),
		Callbacks: make(map[string]telegram.Callback),
	}
}

func (b *TelegramBot) startWebhook() error {
	webhook := ext.WebhookOpts{
		Listen:  b.Config.WebhookListen,
		Port:    b.Config.WebhookPort,
		URLPath: b.Config.WebhookPath + b.Config.BotAPIKey,
	}

	// Delete webhook before starting the bot.
	_, err := b.Bot.DeleteWebhook(&gotgbot.DeleteWebhookOpts{DropPendingUpdates: false})
	if err != nil {
		return err
	}

	err = b.Updater.StartWebhook(b.Bot, webhook)
	if err != nil {
		return err
	}

	_, err = b.Bot.SetWebhook(b.Config.WebhookURL+b.Config.WebhookPath+b.Config.BotAPIKey, &gotgbot.SetWebhookOpts{MaxConnections: 40})
	if err != nil {
		return err
	}

	log.Printf("%s has been started...\n", b.Bot.User.Username)
	return nil
}

func (b *TelegramBot) startPolling() error {
	// Delete webhook before starting the bot
	_, err := b.Bot.DeleteWebhook(&gotgbot.DeleteWebhookOpts{DropPendingUpdates: false})
	if err != nil {
		return err
	}

	err = b.Updater.StartPolling(b.Bot, &ext.PollingOpts{DropPendingUpdates: false})
	if err != nil {
		return err
	}

	log.Printf("%s has been started...\n", b.Bot.User.Username)
	return nil
}

func (b *TelegramBot) startUpdater() error {
	if b.Config.WebhookURL != "" {
		return b.startWebhook()
	} else {
		return b.startPolling()
	}
}

func (b *TelegramBot) Run() (err error) {
	// Connect to DB
	b.newMongo()

	// Create bot from config.
	b.Bot, err = gotgbot.NewBot(b.Config.BotAPIKey, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		return err
	}

	// Create updater and dispatcher.
	b.Updater = ext.NewUpdater(nil)
	b.Dispatcher = b.Updater.Dispatcher

	// Register handlers.
	b.registerHandlers()

	// Load Registered Modules.
	err = b.LoadModules()
	if err != nil {
		return err
	}

	// Start Updating.
	return b.startUpdater()
}
