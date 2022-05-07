package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func UsernameFilter(msg *gotgbot.Message) bool {
	return msg.From.Username == "" && msg.From.Id != 777000
}

func UsernameAndGroupFilter(msg *gotgbot.Message) bool {
	return UsernameFilter(msg) && IsGroup(msg.Chat.Type)
}

func GroupFilter(msg *gotgbot.Message) bool {
	return IsGroup(msg.Chat.Type)
}

func ProfileFilter(bot *gotgbot.Bot, msg *gotgbot.Message) bool {
	p, err := bot.GetUserProfilePhotos(msg.From.Id, &gotgbot.GetUserProfilePhotosOpts{Limit: 1})
	if err != nil {
		return false
	}

	return p.TotalCount == 0
}

func ProfileAndGroupFilter(bot *gotgbot.Bot) func(msg *gotgbot.Message) bool {
	return func(msg *gotgbot.Message) bool {
		return ProfileFilter(bot, msg) && IsGroup(msg.Chat.Type)
	}
}

func AllCallbackFilter(_ *gotgbot.CallbackQuery) bool {
	return true
}
