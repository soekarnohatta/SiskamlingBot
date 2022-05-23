package types

import (
	"SiskamlingBot/bot/core/telegram"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type CallbackFunc = func(*telegram.TgContext) error

type Callback struct {
	Name        string
	Description string
	Callback    string
	Func        CallbackFunc
}

func (cmd Callback) Invoke(bot *gotgbot.Bot, ctx *ext.Context) error {
	newCmdCtx := telegram.NewContext(bot, ctx, "")
	if newCmdCtx != nil {
		return cmd.Func(newCmdCtx)
	}
	return ext.ContinueGroups
}
