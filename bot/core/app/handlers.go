package app

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
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

func (b *MyApp) antispamHandler(bot *gotgbot.Bot, ctx *ext.Context) (ret error) {
	if ctx.Message.NewChatMembers != nil {
		for _, user := range ctx.Message.NewChatMembers {
			if IsBan(user.Id) {
				//log.Printf("User %v is Banned", user.Id)

				dataMap := map[string]string{"1": telegram.MentionHtml(int(user.Id), user.FirstName), "2": utils.IntToStr(int(user.Id))}
				text, keyb := telegram.CreateMenuf("./data/menu/spam.json", 1, dataMap)
				sendOpt := &gotgbot.SendMessageOpts{
					ParseMode:   "HTML",
					ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb},
				}

				_, err := bot.KickChatMember(ctx.Message.Chat.Id, user.Id, nil)
				if err != nil {
					text += " Tetapi saya tidak bisa mengeluarkannya, mohon periksa kembali perizinan saya!"
					_, _ = bot.SendMessage(ctx.Message.Chat.Id, text, sendOpt)
					return ext.EndGroups
				}

				_, _ = bot.SendMessage(ctx.Message.Chat.Id, text, sendOpt)
				return ext.EndGroups
			}
		}
	} else if ctx.Message != nil {
		user := ctx.EffectiveUser
		if IsBan(user.Id) {
			dataMap := map[string]string{"1": telegram.MentionHtml(int(user.Id), user.FirstName), "2": utils.IntToStr(int(user.Id))}
			text, keyb := telegram.CreateMenuf("./data/menu/spam.json", 1, dataMap)
			sendOpt := &gotgbot.SendMessageOpts{
				ParseMode:   "HTML",
				ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb},
			}

			_, err := bot.KickChatMember(ctx.Message.Chat.Id, user.Id, nil)
			if err != nil {
				text += "Tetapi saya tidak bisa mengeluarkannya, mohon periksa kembali perizinan saya!"
				_, _ = bot.SendMessage(ctx.Message.Chat.Id, text, sendOpt)
				return ext.EndGroups
			}

			_, _ = bot.SendMessage(ctx.Message.Chat.Id, text, sendOpt)
			return ext.EndGroups
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

	// Antispam handler
	dsp.AddHandlerToGroup(handlers.NewMessage(filters.All, b.antispamHandler), 1)

	// Other handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(filters.NewChatMembers, b.welcomeHandler), 1)

	// Message handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(filters.NewChatMembers, b.messageHandler), 2)
	dsp.AddHandlerToGroup(handlers.NewMessage(filters.All, b.messageHandler), 2)

	// Callback handlers
	dsp.AddHandlerToGroup(handlers.NewCallback(telegram.AllCallbackFilter, b.callbackHandler), 3)
}
