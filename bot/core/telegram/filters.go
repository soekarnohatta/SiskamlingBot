package telegram

import (
	"regexp"

	"github.com/PaulSonOfLars/gotgbot/v2"
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

func TextCmdPredicate(m *gotgbot.Message) bool {
	return m.Text != "" && m.Text[0] == '/'
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
