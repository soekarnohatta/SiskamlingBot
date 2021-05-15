package core

import (
	"SiskamlingBot/bot/core/telegram"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"log"
	"regexp"
	"strings"
)

/*
 * Group -20: leave channels
 */

func channelPredicate(m *gotgbot.Message) bool {
	return m.Chat.Type == "channel"
}

func (b *TelegramBot) channelLeaveHandler(_ *gotgbot.Bot, ctx *ext.Context) (ret error) {
	// Leave immediately since we already checked the chat type in the filter
	_, err := b.Bot.LeaveChat(ctx.Update.Message.Chat.Id)
	if err != nil {
		log.Println("Leaving channel error")
		return ext.ContinueGroups
	}

	return
}

/*
 * Group 0: command messages
 */

// Run first & propagate to allow telegram to be used in photo captions
func (b *TelegramBot) captionCmdHandler(_ *gotgbot.Bot, ctx *ext.Context) error {
	ctx.Message.Text = ctx.Message.Caption
	return ext.ContinueGroups
}

func textCmdPredicate(m *gotgbot.Message) bool {
	return m.Text != "" && m.Text[0] == '/'
}

func (b *TelegramBot) textCmdHandler(bot *gotgbot.Bot, ctx *ext.Context) (ret error) {
	// Extract only the command invocation segment (first part) from the message text
	// This avoids using strings.Fields() to avoid excess allocations for long, whitespace-heavy messages since
	// we might not even end up invoking a command
	/*
		var cmdSeg string
		cmdEndIdx := strings.IndexFunc(ctx.EffectiveMessage.Text, unicode.IsSpace)
		if cmdEndIdx == -1 {
			cmdSeg = ctx.EffectiveMessage.Text[:util.Min(len(ctx.EffectiveMessage.Text), b.maxCmdSegLen)]
		} else {
			cmdSeg = ctx.EffectiveMessage.Text[:cmdEndIdx]
		}

		// Handle telegram directed towards specific bots with an @username suffix
		unameIdx := strings.IndexByte(cmdSeg, '@')
		if unameIdx != -1 {
			// Extract target username
			uname := cmdSeg[unameIdx:]

			// Ignore command if username doesn't match
			if uname != b.Bot.User.Username {
				return
			}
		}
	*/

	text := ctx.EffectiveMessage.Text
	if ctx.Message.Caption != "" {
		text = ctx.Message.Caption
	}

	var cmd string
	split := strings.Split(strings.ToLower(strings.Fields(text)[0]), "@")
	if len(split) > 1 && split[1] != strings.ToLower(bot.User.Username) {
		cmd = ""
	}

	cmd = split[0][1:]

	// Get and invoke command if valid
	// cmdName := strings.ToLower(cmdSeg[0:])
	if command, ok := b.Commands[cmd]; ok {
		// Invoke command
		command.Invoke(bot, ctx, cmd)
	}

	return ext.ContinueGroups
}

/*
 * Group 1: message listener
 */

func (b *TelegramBot) messageHandler(bot *gotgbot.Bot, ctx *ext.Context) (ret error) {
	if ctx.Message.NewChatMembers != nil || ctx.Message != nil || ctx.Update != nil {
		//TODO: add message filter
		for _, messages := range b.Messages {
			messages.Invoke(bot, ctx)
		}
		return ext.ContinueGroups
	}
	return ext.ContinueGroups
}

/*
 * Group 2: callback listener
 */

func (b *TelegramBot) callbackHandler(bot *gotgbot.Bot, ctx *ext.Context) (ret error) {
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

func (b *TelegramBot) registerHandlers() {
	dsp := b.Dispatcher

	// Channel leave module
	channelHandler := handlers.NewMessage(channelPredicate, b.channelLeaveHandler)
	channelHandler.AllowChannel = true
	dsp.AddHandlerToGroup(channelHandler, -20)

	// Command message handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(filters.Caption, b.captionCmdHandler), 0)
	dsp.AddHandlerToGroup(handlers.NewMessage(textCmdPredicate, b.textCmdHandler), 0)

	// Message handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(filters.NewChatMembers, b.messageHandler), 1)
	dsp.AddHandlerToGroup(handlers.NewMessage(filters.All, b.messageHandler), 1)

	// Callback handlers
	dsp.AddHandlerToGroup(handlers.NewCallback(telegram.AllCallbackFilter, b.callbackHandler), 2)
}
