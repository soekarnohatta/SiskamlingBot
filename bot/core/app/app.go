package app

import (
	"SiskamlingBot/bot/core/telegram"
	"log"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"go.mongodb.org/mongo-driver/mongo"
)

type MyApp struct {
	Bot     *gotgbot.Bot
	Updater ext.Updater
	Context *ext.Context

	Modules   map[string]Module
	Commands  map[string]telegram.Command
	Messages  map[string]telegram.Message
	Callbacks map[string]telegram.Callback

	Config *Config
	DB     *mongo.Database
}

func NewBot(config *Config) *MyApp {
	return &MyApp{
		Context: nil,
		Config:  config,

		Modules:   make(map[string]Module),
		Commands:  make(map[string]telegram.Command),
		Messages:  make(map[string]telegram.Message),
		Callbacks: make(map[string]telegram.Callback),
	}
}

func (b *MyApp) startWebhook() error {
	webhook := ext.WebhookOpts{
		Listen:  b.Config.WebhookListen,
		Port:    b.Config.WebhookPort,
		URLPath: b.Config.WebhookPath + b.Config.BotAPIKey,
	}

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

	log.Printf("%s is now running!\n", b.Bot.User.Username)
	return nil
}

func (b *MyApp) startPolling() error {
	_, err := b.Bot.DeleteWebhook(&gotgbot.DeleteWebhookOpts{DropPendingUpdates: false})
	if err != nil {
		return err
	}

	err = b.Updater.StartPolling(b.Bot, &ext.PollingOpts{DropPendingUpdates: false})
	if err != nil {
		return err
	}

	log.Printf("%s is now running!\n", b.Bot.User.Username)
	return nil
}

func (b *MyApp) startUpdater() error {
	if b.Config.WebhookURL != "" {
		return b.startWebhook()
	} else {
		return b.startPolling()
	}
}

func (b *MyApp) Run() {
	b.newMongo()

	newBotOpt := &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	}

	var err error
	b.Bot, err = gotgbot.NewBot(b.Config.BotAPIKey, newBotOpt)
	if err != nil {
		log.Fatal(err.Error())
	}

	b.Updater = ext.NewUpdater(nil)
	b.registerHandlers()
	b.loadModules()
	b.startUpdater()
}
