package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// IsGroup whether the chat type is Supergroup or not.
func IsGroup(t string) bool {
	return t == "supergroup"
}

// IsPrivate whether the chat type is Private or not.
func IsPrivate(t string) bool {
	return t == "private"
}

// RequireGroup whether the chat type is Supergroup or not.
func RequireGroup(b *gotgbot.Bot, ctx *ext.Context) error {
	if !IsGroup(ctx.Message.Chat.Type) {
		_, err := ctx.Message.Reply(b, "Perintah ini hanya bisa digunakan dalam grup!", nil)
		return err
	}
	return nil
}

// RequirePrivate whether the chat type is Supergroup or not.
func RequirePrivate(b *gotgbot.Bot, ctx *ext.Context) error {
	if !IsPrivate(ctx.Message.Chat.Type) {
		_, err := ctx.Message.Reply(b, "Perintah ini hanya bisa digunakan dalam japri!", nil)
		return err
	}
	return nil
}

