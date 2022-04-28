package app

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/core/telegram/types"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

/*
 * Group 0: command messages
 */

func (*MyApp) captionCmdHandler(_ *gotgbot.Bot, ctx *ext.Context) error {
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
	if len(split) > 1 && strings.ToLower(bot.User.Username) != split[1] {
		return ext.ContinueGroups
	}

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
		var wg sync.WaitGroup
		var orderedGroup []types.Message

		for _, z := range b.Modules {
			for _, y := range z.Messages() {
				orderedGroup = append(orderedGroup, y)
			}
		}

		sort.SliceStable(orderedGroup, func(i, j int) bool {
			return orderedGroup[i].Order < orderedGroup[j].Order
		})

		for _, messages := range orderedGroup {
			if messages.Filter == nil {
				messages.Filter = message.All
			}

			if messages.Filter(ctx.Message) {
				if messages.Async {
					wg.Add(1)
					go messages.InvokeAsync(&wg, bot, ctx)
				} else {
					messages.Invoke(bot, ctx)
				}
			}
		}

		wg.Wait()
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

/*
 * Group -10: bot misc
 */

func (b *MyApp) welcomeHandler(bot *gotgbot.Bot, ctx *ext.Context) (ret error) {
	if ctx.Message.NewChatMembers != nil {
		for _, user := range ctx.Message.NewChatMembers {
			if user.Id == b.Bot.User.Id {
				dataMap := map[string]string{"1": b.Bot.User.FirstName, "2": b.Config.BotVer, "3": "Unknown", "uname": b.Bot.User.Username}
				text, keyb := telegram.CreateMenuf("./data/menu/start.json", 2, dataMap)
				sendOpt := &gotgbot.SendMessageOpts{
					ParseMode:   "HTML",
					ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb},
				}
				_, _ = bot.SendMessage(ctx.Message.Chat.Id, text, sendOpt)
			}
		}
		return ext.ContinueGroups
	}
	return ext.ContinueGroups
}

func (b *MyApp) registerHandlers() {
	dsp := b.Updater.Dispatcher

	// Command message handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(message.Caption, b.captionCmdHandler), 0)
	dsp.AddHandlerToGroup(handlers.NewMessage(telegram.TextCmdPredicate, b.textCmdHandler), 0)

	// Callback handlers
	dsp.AddHandlerToGroup(handlers.NewCallback(telegram.AllCallbackFilter, b.callbackHandler), 1)

	// Other handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(message.NewChatMembers, b.welcomeHandler), 2)

	// Message handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(message.NewChatMembers, b.messageHandler), 3)
	dsp.AddHandlerToGroup(handlers.NewMessage(message.All, b.messageHandler), 3)
}
