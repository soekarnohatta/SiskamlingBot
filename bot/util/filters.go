package util

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"regexp"
)

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

func AllCallbackFilter(_ *gotgbot.CallbackQuery) bool {
	return true
}

func CallbackRegexFilter(expr string) func(cq *gotgbot.CallbackQuery) bool {
	return func(cq *gotgbot.CallbackQuery) bool {
		pattern, _ := regexp.Compile(expr)
		return pattern.MatchString(cq.Data)
	}
}
