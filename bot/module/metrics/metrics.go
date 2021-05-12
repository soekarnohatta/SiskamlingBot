package metrics

import (
	"SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
)

// Module contains the state for an instance of this module.
type Module struct {
	Bot *core.TelegramBot
}

// Info returns basic information about this module.
func (m *Module) Info() core.ModuleInfo {
	return core.ModuleInfo{
		Name: "Metrics",
	}
}

// Commands returns a list of telegram provided by this module.
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
			Func:        m.chatMetric,
		},
		{
			Name:        "userMetric",
			Description: "Detect user without profile picture",
			Func:        m.usernameMetric,
		},
	}
}

func (m *Module) Callbacks() []telegram.Callback {
	return []telegram.Callback{}
}

// NewModule returns a new instance of this module.
func NewModule(bot *core.TelegramBot) (core.Module, error) {
	return &Module{
		Bot: bot,
	}, nil
}

func init() {
	core.RegisterModule("Metrics", NewModule)
}
