package app

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"

	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/core/telegram/types"
)

func (b *MyApp) registerCommandUsingDispatcher() {
	defer b.handlePanicSendLog(nil)
	for _, cmd := range b.Commands {
		b.Updater.Dispatcher.AddHandler(newCustomCommandHandler(cmd.Trigger, cmd.InvokeWithDispatcher))
	}
}

func (b *MyApp) messageHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	defer b.handlePanicSendLog(ctx)
	if ctx.EffectiveMessage == nil {
		return ext.ContinueGroups
	}

	var orderedGroup []types.Message

	for _, y := range b.Messages {
		orderedGroup = append(orderedGroup, y)
	}

	sort.SliceStable(orderedGroup, func(i, j int) bool {
		return orderedGroup[i].Order < orderedGroup[j].Order
	})

	var wg sync.WaitGroup
	for _, messages := range orderedGroup {
		if messages.Filter == nil {
			messages.Filter = message.All
		}

		if messages.Filter(ctx.EffectiveMessage) {
			if messages.Async == true {
				wg.Add(1)
				go func(wg *sync.WaitGroup, bot *gotgbot.Bot, ctx *ext.Context) {
					defer wg.Done()
					defer b.handlePanicSendLog(ctx)

					// we use our own message handler, we process update manually so we can't use
					// the error handling middleware provided from the library
					err := messages.InvokeAsync(bot, ctx)
					if errors.Is(err, telegram.EndOrder) {
						return
					} else if errors.Is(err, telegram.ContinueOrder) {
						return
					} else if err != nil {
						b.SendLogMessage("Error Async Message Handler", err, ctx)
						return
					}
				}(&wg, bot, ctx)
			} else {
				// we use our own message handler, we process update manually so we can't use
				// the error handling middleware provided from the library
				err := messages.Invoke(bot, ctx)
				if errors.Is(err, telegram.EndOrder) {
					return nil
				} else if errors.Is(err, telegram.ContinueOrder) || err.Error() == telegram.ContinueOrder.Error() {
					continue
				} else if err != nil {
					b.SendLogMessage("Error Message Handler", err, ctx)
					continue
				}
			}
		}
		wg.Wait()
	}
	return ext.ContinueGroups
}

func (b *MyApp) callbackHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	defer b.handlePanicSendLog(ctx)
	if ctx.CallbackQuery != nil || ctx.Update.CallbackQuery != nil {
		for _, callbacks := range b.Callbacks {
			pattern, _ := regexp.Compile(callbacks.Callback)
			if pattern.MatchString(ctx.CallbackQuery.Data) {
				err := callbacks.Invoke(bot, ctx)
				if errors.Is(err, telegram.EndOrder) {
					return nil
				} else if errors.Is(err, telegram.ContinueOrder) ||
					(err != nil && err.Error() == telegram.ContinueOrder.Error()) {
					continue
				} else if err != nil {
					b.SendLogMessage("Error Callback Handler", err, ctx)
					continue
				}
			}
		}
	}
	return ext.ContinueGroups
}

func (b *MyApp) handlePanicSendLog(ctx *ext.Context) {
	if r := recover(); r != nil {
		b.SendLogMessage("Recover Panic Error", fmt.Errorf("%v", r), ctx)
	}
}

func newCustomMessageHandler(f filters.Message, r handlers.Response) handlers.Message {
	return handlers.Message{
		AllowEdited:  true,
		AllowChannel: false,
		Filter:       f,
		Response:     r,
	}
}

func newCustomCommandHandler(cmd string, r handlers.Response) handlers.Command {
	return handlers.Command{
		Triggers:     []rune{'/', '!', ','},
		AllowEdited:  true,
		AllowChannel: false,
		Command:      cmd,
		Response:     r,
	}
}

func (b *MyApp) registerHandlers() {
	dsp := b.Updater.Dispatcher

	// Command message handlers
	b.registerCommandUsingDispatcher()

	// Callback handlers
	dsp.AddHandlerToGroup(handlers.NewCallback(telegram.AllCallbackFilter, b.callbackHandler), 1)

	// Message handlers
	dsp.AddHandlerToGroup(handlers.NewMessage(message.NewChatMembers, b.messageHandler), 2)
	dsp.AddHandlerToGroup(newCustomMessageHandler(message.All, b.messageHandler), 2)

	b.ErrorLog.Println("All handlers have been registered successfully!")
}
