package command

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Ping(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.Message.Reply(b, "Ping", nil)
	return err
}
