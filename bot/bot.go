package bot

import (
	"SiskamlingBot/bot/handler/command"
	"SiskamlingBot/bot/handler/listener"
	"SiskamlingBot/bot/handler/listener/callbacks"
	"SiskamlingBot/bot/handler/listener/metrics"
	"SiskamlingBot/bot/helper/telegram"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

func AddHandler(bot *gotgbot.Bot, dispatcher *ext.Dispatcher) {
	dispatcher.AddHandler(handlers.NewCommand("ping", command.Ping))

	dispatcher.AddHandlerToGroup(handlers.NewMessage(filters.All, metrics.ChatMetrics), 0)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(filters.All, metrics.UsernameMetrics), 0)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(telegram.UsernameAndGroupFilter, listener.Username), 0)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(telegram.ProfileAndGroupFilter(bot), listener.Picture), 0)

	dispatcher.AddHandlerToGroup(handlers.NewCallback(filters.Prefix("username("), callbacks.UsernameCB), 1)
	dispatcher.AddHandlerToGroup(handlers.NewCallback(filters.Prefix("picture("), callbacks.PictureCB), 1)
}
