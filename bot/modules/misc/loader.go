package user

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
		Name: "Misc",
	}
}

// Commands returns a list of telegram provided by this module.
func (m Module) Commands() []telegram.Command {
	return []telegram.Command{
		{
			Name:        "ping",
			Description: "ping the bot.",
			Func:        m.ping,
		},
		{
			Name:        "about",
			Description: "about the bot.",
			Func:        m.about,
		},
		{
			Name:        "start",
			Description: "start the bot.",
			Func:        m.start,
		},
	}
}

func (m Module) Messages() []telegram.Message {
	return nil
}

func (m Module) Callbacks() []telegram.Callback {
	return nil
}

// NewModule returns a new instance of this module.
func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{
		App: bot,
	}, nil
}

func init() {
	app.RegisterModule("Misc", NewModule)
}
