package telegram

import (
	"sync"

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
	Async		bool
}

func (cmd Message) Invoke(bot *gotgbot.Bot, ctx *ext.Context) {
	newCmdCtx := NewContext(bot, ctx, "")
	if newCmdCtx != nil {
		cmd.Func(newCmdCtx)
	}
}

func (cmd Message) InvokeAsync(wg *sync.WaitGroup, bot *gotgbot.Bot, ctx *ext.Context) {
	defer wg.Done()
	newCmdCtx := NewContext(bot, ctx, "")
	if newCmdCtx != nil {
		cmd.Func(newCmdCtx)
	}
}
