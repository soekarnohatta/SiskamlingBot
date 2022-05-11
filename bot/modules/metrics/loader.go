package metrics

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram"
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
			Name:   "Chat Metric",
			Filter: telegram.GroupFilter,
			Func:   m.chatMetric,
			Order:  0,
			Async:  true,
		},
		{
			Name:   "User Metric",
			Filter: message.All,
			Func:   m.usernameMetric,
			Order:  0,
			Async:  true,
		},
		{
			Name:   "Prefence Metric",
			Filter: telegram.GroupFilter,
			Func:   m.preferenceMetric,
			Order:  0,
			Async:  true,
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
	err := app.RegisterModule("Metrics", NewModule)
	if err != nil {
		panic(err)
	}
}
