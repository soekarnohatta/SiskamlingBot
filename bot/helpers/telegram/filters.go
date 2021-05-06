package telegram

import "github.com/PaulSonOfLars/gotgbot/v2"

func UsernameFilter(msg *gotgbot.Message) bool {
	return msg.From.Username == "" && msg.From.Id != 777000
}

func UsernameAndGroupFilter(msg *gotgbot.Message) bool {
	return UsernameFilter(msg) && IsGroup(msg.Chat.Type)
}

func ProfileFilter(bot *gotgbot.Bot, msg *gotgbot.Message) bool {
	p, err := msg.From.GetProfilePhotos(bot, nil)
	return err == nil && p != nil && p.TotalCount == 0
}

func ProfileAndGroupFilter(bot *gotgbot.Bot) func(msg *gotgbot.Message) bool {
	return func(msg *gotgbot.Message) bool {
		return ProfileFilter(bot, msg) && IsGroup(msg.Chat.Type)
	}
}
