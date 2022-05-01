package app

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/core/telegram/types"
	"errors"
	"regexp"
	"sort"
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

func (b *MyApp) registerCommandUsingDispatcher() {
	for _, cmd := range b.Commands {
		b.Updater.Dispatcher.AddHandler(&handlers.Command{
			Triggers:     []rune{'/', '!'},
			AllowEdited:  false,
			AllowChannel: false,
			Command:      cmd.Trigger,
			Response:     cmd.InvokeWithDispatcher,
		})
	}
}

func (b *MyApp) messageHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.Message.NewChatMembers != nil || ctx.Message != nil || ctx.Update.EditedMessage != nil || ctx.Update != nil {
		var orderedGroup []types.Message

		for _, y := range b.Messages {
			orderedGroup = append(orderedGroup, y)
		}

		sort.SliceStable(orderedGroup, func(i, j int) bool {
			return orderedGroup[i].Order < orderedGroup[j].Order
		})

		limit := make(chan struct{}, 5)
		var wg sync.WaitGroup
		for _, messages := range orderedGroup {
			if messages.Filter == nil {
				messages.Filter = message.All
			}

			if messages.Filter(ctx.Message) {
				if messages.Async == true {
					wg.Add(1)
					limit <- struct{}{}
					go func(wg *sync.WaitGroup, bot *gotgbot.Bot, ctx *ext.Context) {
						defer wg.Done()
						// TODO: Add log to group
						_ = messages.InvokeAsync(bot, ctx)
						<-limit
					}(&wg, bot, ctx)
				} else {
					err := messages.Invoke(bot, ctx)
					if errors.Is(err, telegram.EndOrder) {
						return nil
					} else if errors.Is(err, telegram.ContinueOrder) {
						continue
					} else {
						b.SendLogMessage("Error", err)
						continue
					}
				}
			}
			wg.Wait()
		}
		return ext.ContinueGroups
	}
	return ext.ContinueGroups
}

func (b *MyApp) callbackHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.CallbackQuery != nil || ctx.Update.CallbackQuery != nil {
		for _, callbacks := range b.Callbacks {
			pattern, _ := regexp.Compile(callbacks.Callback)
			if pattern.MatchString(ctx.CallbackQuery.Data) {
				err := callbacks.Invoke(bot, ctx)
				if err != nil {
					return err
				}
			}
		}
	}
	return ext.ContinueGroups
}

func (b *MyApp) registerHandlers() {
	dsp := b.Updater.Dispatcher

	// Command message handlers
	b.registerCommandUsingDispatcher()

	// Callback handlers
	dsp.AddHandlerToGroup(handlers.NewCallback(telegram.AllCallbackFilter, b.callbackHandler), 1)

	// Message handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(message.NewChatMembers, b.messageHandler), 2)
	dsp.AddHandlerToGroup(handlers.NewMessage(message.All, b.messageHandler), 2)

	b.ErrorLog.Println("All handlers have been registered succesfully!")
}
