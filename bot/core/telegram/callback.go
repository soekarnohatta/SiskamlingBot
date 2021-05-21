package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type CallbackFunc = func(*TgContext)

type Callback struct {
	Name        string
	Description string
	Callback    string
	Func        CallbackFunc
}

func (cmd Callback) Invoke(bot *gotgbot.Bot, ctx *ext.Context) {
	newCmdCtx := newContext(bot, ctx, "")
	cmd.Func(newCmdCtx)
}
