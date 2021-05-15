package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// MessageFunc represents a command function that takes no message arguments.
type MessageFunc = func(*TgContext)

// Message describes a bot command.
type Message struct {
	Name        string
	Description string
	Func        MessageFunc
}

// Invoke invokes a Message with the given arguments.
func (cmd Message) Invoke(bot *gotgbot.Bot, ctx *ext.Context) {
	// Construct context
	newCmdCtx := newContext(bot, ctx, "")
	cmd.Func(newCmdCtx)
}
