package types

import (
	"SiskamlingBot/bot/core/telegram"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

type MessageFunc = func(*telegram.TgContext) error

type Message struct {
	Name        string
	Description string
	Filter      filters.Message
	Func        MessageFunc
	Order       int
	Async       bool
}

func (cmd Message) Invoke(bot *gotgbot.Bot, ctx *ext.Context) error {
	newCmdCtx := telegram.NewContext(bot, ctx, "")
	if newCmdCtx != nil {
		return cmd.Func(newCmdCtx)
	}
	return ext.ContinueGroups
}

func (cmd Message) InvokeAsync(bot *gotgbot.Bot, ctx *ext.Context) error {
	newCmdCtx := telegram.NewContext(bot, ctx, "")
	if newCmdCtx != nil && cmd.Async {
		return cmd.Func(newCmdCtx)
	}
	return ext.ContinueGroups
}
