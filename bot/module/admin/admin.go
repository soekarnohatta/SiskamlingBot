package admin

import (
	"SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
)

// Module contains the state for an instance of this module.
type Module struct {
	App *core.MyApp
}

// Info returns basic information about this module.
func (m *Module) Info() core.ModuleInfo {
	return core.ModuleInfo{
		Name: "Admin",
	}
}

// Commands returns a list of telegram provided by this module.
func (m *Module) Commands() []telegram.Command {
	return []telegram.Command{
		{
			Name:        "user",
			Description: "get user info",
			Func:        m.getUser,
		},
		{
			Name:        "chat",
			Description: "get chat info",
			Func:        m.getChat,
		},
		{
			Name:        "dbg",
			Description: "debug",
			Func:        m.debug,
		},
	}
}

func (m *Module) Messages() []telegram.Message {
	return nil
}

func (m *Module) Callbacks() []telegram.Callback {
	return nil
}

// NewModule returns a new instance of this module.
func NewModule(bot *core.MyApp) (core.Module, error) {
	return &Module{
		App: bot,
	}, nil
}

func init() {
	core.RegisterModule("Admin", NewModule)
}
