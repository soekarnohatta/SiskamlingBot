package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// CommandFunc represents a command function that takes no message arguments.
type CommandFunc = func(*TgContext)

// Command describes a bot command.
type Command struct {
	Name        string
	Description string
	Usage       string
	Aliases     []string
	Func        CommandFunc
}

// Invoke invokes a Command with the given arguments.
func (cmd *Command) Invoke(bot *gotgbot.Bot, ctx *ext.Context, cmdSeg string) {
	// Construct context
	newCmdCtx := newContext(bot, ctx, cmdSeg)
	cmd.Func(newCmdCtx)
}
