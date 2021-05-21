package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

type MessageFunc = func(*TgContext)

type Message struct {
	Name        string
	Description string
	Filter      filters.Message
	Func        MessageFunc
}

func (cmd Message) Invoke(bot *gotgbot.Bot, ctx *ext.Context) {
	newCmdCtx := newContext(bot, ctx, "")
	if newCmdCtx != nil {
		cmd.Func(newCmdCtx)
	}
}
