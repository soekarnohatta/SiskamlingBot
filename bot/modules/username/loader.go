package username

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram"
)

// Module contains the state for an instance of this module.
type Module struct {
	App *app.MyApp
}

// Info returns basic information about this module.
func (m Module) Info() app.ModuleInfo {
	return app.ModuleInfo{
		Name: "Username",
	}
}

// Commands returns a list of telegram provided by this module.
func (m Module) Commands() []telegram.Command {
	return []telegram.Command{
	}
}

func (m Module) Messages() []telegram.Message {
	return []telegram.Message{
		{
			Name:        "UsernameScanner",
			Description: "Detect user without username",
			Filter:      telegram.UsernameAndGroupFilter,
			Func:        m.usernameScan,
			Async: 		 false,
		},
	}
}

func (m Module) Callbacks() []telegram.Callback {
	return []telegram.Callback{
		{
			Name:        "UsernameCallback",
			Description: "",
			Callback:    `username\((.+?)\)`,
			Func:        m.usernameCallback,
		},
	}
}

// NewModule returns a new instance of this module.
func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{
		App: bot,
	}, nil
}

func init() {
	app.RegisterModule("Username", NewModule)
}
