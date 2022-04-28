package types

import (
	"SiskamlingBot/bot/core/telegram"
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

type MessageFunc = func(*telegram.TgContext)

type Message struct {
	Name        string
	Description string
	Filter      filters.Message
	Func        MessageFunc
	Order       int
	Async       bool
}

func (cmd Message) Invoke(bot *gotgbot.Bot, ctx *ext.Context) {
	newCmdCtx := telegram.NewContext(bot, ctx, "")
	if newCmdCtx != nil {
		cmd.Func(newCmdCtx)
	}
}

func (cmd Message) InvokeAsync(wg *sync.WaitGroup, bot *gotgbot.Bot, ctx *ext.Context) {
	defer wg.Done()
	newCmdCtx := telegram.NewContext(bot, ctx, "")
	if newCmdCtx != nil {
		cmd.Func(newCmdCtx)
	}
}
