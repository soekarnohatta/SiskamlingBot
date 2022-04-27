package metrics

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

type Module struct {
	App *app.MyApp
}

func (m *Module) Info() app.ModuleInfo {
	return app.ModuleInfo{
		Name: "Metrics",
	}
}

func (m *Module) Commands() []telegram.Command {
	return nil
}

func (m *Module) Messages() []telegram.Message {
	return []telegram.Message{
		{
			Name:        "chatMetric",
			Description: "Detect user without username",
			Filter:      message.All,
			Func:        m.chatMetric,
			Async:       true,
		},
		{
			Name:        "userMetric",
			Description: "Detect user without profile picture",
			Filter:      message.All,
			Func:        m.usernameMetric,
			Async:       true,
		},
	}
}

func (m *Module) Callbacks() []telegram.Callback {
	return nil
}

func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{
		App: bot,
	}, nil
}

func init() {
	app.RegisterModule("Metrics", NewModule)
}
