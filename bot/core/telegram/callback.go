package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// CallbackFunc represents a command function that takes no message arguments.
type CallbackFunc = func(*TgContext)

// Callback describes a bot command.
type Callback struct {
	Name        string
	Description string
	Callback    string
	Func        CallbackFunc
}

// Invoke invokes a Callback with the given arguments.
func (cmd Callback) Invoke(bot *gotgbot.Bot, ctx *ext.Context) {
	// Construct context
	newCmdCtx := newContext(bot, ctx, "")
	cmd.Func(newCmdCtx)
}
