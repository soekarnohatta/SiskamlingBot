package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type CommandFunc = func(*TgContext)

type Command struct {
	Name        string
	Description string
	Usage       string
	Aliases     []string
	Func        CommandFunc
}

func (cmd Command) Invoke(bot *gotgbot.Bot, ctx *ext.Context, cmdSeg string) {
	newCmdCtx := NewContext(bot, ctx, cmdSeg)
	if newCmdCtx != nil {
		cmd.Func(newCmdCtx)
	}
}
