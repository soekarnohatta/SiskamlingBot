package app

import (
	"SiskamlingBot/bot/core/telegram/types"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/shirou/gopsutil/host"
)

type MyApp struct {
	Bot     *gotgbot.Bot
	Updater ext.Updater
	Context *ext.Context

	Modules   map[string]Module
	Commands  map[string]types.Command
	Messages  map[string]types.Message
	Callbacks map[string]types.Callback

	Config   *Config
	DB       MongoDB
	ErrorLog *log.Logger
}

func NewBot(config *Config) *MyApp {
	return &MyApp{
		Context:  nil,
		Config:   config,
		ErrorLog: log.New(os.Stderr, "[BOT]", log.LstdFlags),

		Modules:   make(map[string]Module),
		Commands:  make(map[string]types.Command),
		Messages:  make(map[string]types.Message),
		Callbacks: make(map[string]types.Callback),
	}
}

func (b *MyApp) startWebhook() error {
	webhook := ext.WebhookOpts{
		Listen:  b.Config.WebhookListen,
		Port:    b.Config.WebhookPort,
		URLPath: b.Config.WebhookPath + b.Config.BotAPIKey,
	}

	_, err := b.Bot.DeleteWebhook(&gotgbot.DeleteWebhookOpts{DropPendingUpdates: b.Config.CleanPolling})
	if err != nil {
		return err
	}

	err = b.Updater.StartWebhook(b.Bot, webhook)
	if err != nil {
		return err
	}

	_, err = b.Bot.SetWebhook(b.Config.WebhookURL+b.Config.WebhookPath+b.Config.BotAPIKey, &gotgbot.SetWebhookOpts{
		MaxConnections:     40,
		DropPendingUpdates: b.Config.CleanPolling,
	})
	if err != nil {
		return err
	}

	log.Printf("%s is now running using webhook!\n", b.Bot.User.Username)
	return nil
}

func (b *MyApp) startPolling() error {
	_, err := b.Bot.DeleteWebhook(&gotgbot.DeleteWebhookOpts{DropPendingUpdates: b.Config.CleanPolling})
	if err != nil {
		return err
	}

	err = b.Updater.StartPolling(b.Bot, &ext.PollingOpts{DropPendingUpdates: b.Config.CleanPolling})
	if err != nil {
		return err
	}

	b.ErrorLog.Printf("%s is now running using long-polling!\n", b.Bot.User.Username)
	return nil
}

func (b *MyApp) startUpdater() error {
	if b.Config.WebhookURL != "" {
		return b.startWebhook()
	} else {
		return b.startPolling()
	}
}

func (b *MyApp) Run() error {
	newBotOpt := &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	}

	var err error
	b.Bot, err = gotgbot.NewBot(b.Config.BotAPIKey, newBotOpt)
	if err != nil {
		return err
	}

	b.Updater = ext.NewUpdater(nil)

	err = b.newMongo()
	if err != nil {
		return err
	}

	err = b.loadModules()
	if err != nil {
		return err
	}

	b.registerHandlers()

	err = b.startUpdater()
	if err != nil {
		return err
	}

	return nil
}

func (b *MyApp) SendLogMessage(msg string, err error) {
	bot := b.Bot
	info, _ := host.Info()
	replyTxt := fmt.Sprintf("<b>System Log Message</b>\n"+
		"<b>%v</b>\n\n"+
		"<b>Bot Name :</b> %v\n"+
		"<b>Bot Username :</b> @%v\n"+
		"<b>Host OS :</b> %v\n"+
		"<b>Host Name :</b> %v\n"+
		"<b>Host Uptime :</b> %v\n"+
		"<b>Kernel Version :</b> %v\n"+
		"<b>Platform :</b> %v\n"+
		"<b>Timestamp :</b> %v\n",
		msg,
		bot.FirstName,
		bot.Username,
		info.OS,
		info.Hostname,
		info.Uptime,
		info.KernelVersion,
		info.Platform,
		time.Now().Local(),
	)
	if err != nil {
		replyTxt += "=====================\n"
		replyTxt += "<b>Error Details:</b>\n"
		replyTxt += html.EscapeString(err.Error())
	}

	_, err = b.Bot.SendMessage(b.Config.LogEvent, replyTxt, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		b.ErrorLog.Println(err)
	}
}
