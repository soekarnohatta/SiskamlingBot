package metrics

import (
	app "SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
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
	return []telegram.Command{
		{},
	}
}

func (m *Module) Messages() []telegram.Message {
	return []telegram.Message{
		{
			Name:        "chatMetric",
			Description: "Detect user without username",
			Filter:      filters.All,
			Func:        m.chatMetric,
		},
		{
			Name:        "userMetric",
			Description: "Detect user without profile picture",
			Filter:      filters.All,
			Func:        m.usernameMetric,
		},
	}
}

func (m *Module) Callbacks() []telegram.Callback {
	return []telegram.Callback{}
}

func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{
		App: bot,
	}, nil
}

func init() {
	app.RegisterModule("Metrics", NewModule)
}
