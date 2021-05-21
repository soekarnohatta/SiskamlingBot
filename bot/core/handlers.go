package core

import (
	"SiskamlingBot/bot/core/telegram"
	"regexp"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

/*
 * Group 0: command messages
 */

func (b *MyApp) captionCmdHandler(_ *gotgbot.Bot, ctx *ext.Context) error {
	ctx.Message.Text = ctx.Message.Caption
	return ext.ContinueGroups
}

func (b *MyApp) textCmdHandler(bot *gotgbot.Bot, ctx *ext.Context) (ret error) {
	text := ctx.EffectiveMessage.Text
	if ctx.Message.Caption != "" {
		text = ctx.Message.Caption
	}

	var cmd string
	split := strings.Split(strings.ToLower(strings.Fields(text)[0]), "@")
	cmd = split[0][1:]
	if command, ok := b.Commands[cmd]; ok {
		command.Invoke(bot, ctx, cmd)
	}

	return ext.ContinueGroups
}

/*
 * Group 1: message listener
 */

func (b *MyApp) messageHandler(bot *gotgbot.Bot, ctx *ext.Context) (ret error) {
	if ctx.Message.NewChatMembers != nil || ctx.Message != nil || ctx.Update != nil {
		for _, messages := range b.Messages {
			if messages.Filter == nil {
				messages.Filter = filters.All
			}

			if messages.Filter(ctx.Message) {
				messages.Invoke(bot, ctx)
			}
		}
		return ext.ContinueGroups
	}
	return ext.ContinueGroups
}

/*
 * Group 2: callback listener
 */

func (b *MyApp) callbackHandler(bot *gotgbot.Bot, ctx *ext.Context) (ret error) {
	if ctx.CallbackQuery != nil || ctx.Update.CallbackQuery != nil {
		for _, callbacks := range b.Callbacks {
			pattern, _ := regexp.Compile(callbacks.Callback)
			if pattern.MatchString(ctx.CallbackQuery.Data) {
				callbacks.Invoke(bot, ctx)
			}
		}
		return ext.ContinueGroups
	}
	return ext.ContinueGroups
}

func (b *MyApp) registerHandlers() {
	dsp := b.Updater.Dispatcher

	// Command message handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(filters.Caption, b.captionCmdHandler), 0)
	dsp.AddHandlerToGroup(handlers.NewMessage(telegram.TextCmdPredicate, b.textCmdHandler), 0)

	// Message handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(filters.NewChatMembers, b.messageHandler), 1)
	dsp.AddHandlerToGroup(handlers.NewMessage(filters.All, b.messageHandler), 1)

	// Callback handlers
	dsp.AddHandlerToGroup(handlers.NewCallback(telegram.AllCallbackFilter, b.callbackHandler), 2)
}
