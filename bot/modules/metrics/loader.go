package metrics

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram/types"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

type Module struct {
	App *app.MyApp
}

func (*Module) Info() app.ModuleInfo {
	return app.ModuleInfo{Name: "Metrics"}
}

func (*Module) Commands() []types.Command {
	return nil
}

func (m *Module) Messages() []types.Message {
	return []types.Message{
		{
			Name:        "Chat Metric",
			Description: "Saves chat info",
			Filter:      message.All,
			Func:        m.chatMetric,
			Order:       0,
			Async:       true,
		},
		{
			Name:        "User Metric",
			Description: "Saves user info",
			Filter:      message.All,
			Func:        m.usernameMetric,
			Order:       0,
			Async:       true,
		},
	}
}

func (*Module) Callbacks() []types.Callback {
	return nil
}

func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{App: bot}, nil
}

func init() {
	app.RegisterModule("Metrics", NewModule)
}
