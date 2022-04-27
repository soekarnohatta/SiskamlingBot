package app

import "github.com/PaulSonOfLars/gotgbot/v2"

func (b *MyApp) SendLog(text string) {
	sendOpt := &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	}
	_, _ = b.Bot.SendMessage(b.Config.LogEvent, text, sendOpt)
}
