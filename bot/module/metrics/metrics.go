package metrics

import (
	"SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

type Module struct {
	App *core.MyApp
}

func (m *Module) Info() core.ModuleInfo {
	return core.ModuleInfo{
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
			Filter: 	 filters.All,
			Func:        m.chatMetric,
		},
		{
			Name:        "userMetric",
			Description: "Detect user without profile picture",
			Filter: 	 filters.All,
			Func:        m.usernameMetric,
		},
	}
}

func (m *Module) Callbacks() []telegram.Callback {
	return []telegram.Callback{}
}

func NewModule(bot *core.MyApp) (core.Module, error) {
	return &Module{
		App: bot,
	}, nil
}

func init() {
	core.RegisterModule("Metrics", NewModule)
}
