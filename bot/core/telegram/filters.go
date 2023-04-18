package telegram

import (
	"SiskamlingBot/bot/utils"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func UsernameFilter(msg *gotgbot.Message) bool {
	return msg.From.Username == "" && msg.From.Id != 777000
}

func UsernameAndGroupFilter(msg *gotgbot.Message) bool {
	return UsernameFilter(msg) && utils.IsGroup(msg.Chat.Type)
}

func GroupFilter(msg *gotgbot.Message) bool {
	return utils.IsGroup(msg.Chat.Type)
}

func ProfileFilter(bot *gotgbot.Bot, msg *gotgbot.Message) bool {
	p, err := bot.GetUserProfilePhotos(msg.From.Id, &gotgbot.GetUserProfilePhotosOpts{Limit: 1})
	if err != nil {
		return false
	}

	return len(p.Photos) < 1 || p.TotalCount < 1
}

func ProfileAndGroupFilter(bot *gotgbot.Bot) func(msg *gotgbot.Message) bool {
	return func(msg *gotgbot.Message) bool {
		return ProfileFilter(bot, msg) && utils.IsGroup(msg.Chat.Type)
	}
}

func AllCallbackFilter(_ *gotgbot.CallbackQuery) bool {
	return true
}
