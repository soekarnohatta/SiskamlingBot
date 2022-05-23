package types

import (
	"SiskamlingBot/bot/core/telegram"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type CommandFunc = func(*telegram.TgContext) error

type Command struct {
	Name        string
	Trigger     string
	Description string
	Usage       string
	Aliases     []string
	Func        CommandFunc
}

func (cmd Command) Invoke(bot *gotgbot.Bot, ctx *ext.Context, cmdSeg string) error {
	newCmdCtx := telegram.NewContext(bot, ctx, cmdSeg)
	if newCmdCtx != nil {
		return cmd.Func(newCmdCtx)
	}
	return nil
}

func (cmd Command) InvokeWithDispatcher(bot *gotgbot.Bot, ctx *ext.Context) error {
	newCmdCtx := telegram.NewContext(bot, ctx, "")
	if newCmdCtx != nil {
		return cmd.Func(newCmdCtx)
	}
	return nil
}
