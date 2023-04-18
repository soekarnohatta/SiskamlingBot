package utils

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func IsGroup(t string) bool {
	return t == "supergroup" || t == "group"
}

func IsPrivate(t string) bool {
	return t == "private"
}

func IsSudo(u int64, sudo []int64) bool {
	for _, val := range sudo {
		if u == val {
			return true
		}
		continue
	}

	return false
}

func IsOwner(u, owner int64) bool {
	return u == owner
}

func RequireGroup(b *gotgbot.Bot, ctx *ext.Context) error {
	if !IsGroup(ctx.Message.Chat.Type) {
		_, err := ctx.Message.Reply(b, "Perintah ini hanya bisa digunakan dalam grup!", nil)
		return err
	}
	return nil
}

func RequirePrivate(b *gotgbot.Bot, ctx *ext.Context) error {
	if !IsPrivate(ctx.Message.Chat.Type) {
		_, err := ctx.Message.Reply(b, "Perintah ini hanya bisa digunakan dalam japri!", nil)
		return err
	}
	return nil
}

func IsUserAdmin(b *gotgbot.Bot, chatId, userId int64) bool {
	member, err := b.GetChatMember(chatId, userId, nil)
	if err != nil {
		return false
	}

	return member.MergeChatMember().Status == "administrator" || member.MergeChatMember().Status == "creator"
}
