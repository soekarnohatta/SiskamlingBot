package username

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/core/telegram/types"
)

type Module struct {
	App *app.MyApp
}

func (*Module) Info() app.ModuleInfo {
	return app.ModuleInfo{Name: "Username"}
}

func (m *Module) Commands() []types.Command {
	return []types.Command{
		{
			Name:    "setusername",
			Trigger: "setusername",
			Func:    m.usernameSetting,
		},
	}
}

func (m *Module) Messages() []types.Message {
	return []types.Message{
		{
			Name:   "UsernameScanner",
			Filter: telegram.UsernameAndGroupFilter,
			Func:   m.usernameScan,
			Order:  2,
			Async:  false,
		},
	}
}

func (m *Module) Callbacks() []types.Callback {
	return []types.Callback{
		{
			Name:     "UsernameCallbackGroup",
			Callback: `username\((.+?)\)\((.+?)\)`,
			Func:     m.usernameCallbackGroup,
		},
		{
			Name:     "UsernameCallbackPrivate",
			Callback: `username\((.+?)\)\((.+?)\)`,
			Func:     m.usernameCallbackPrivate,
		},
	}
}

// NewModule returns a new instance of this module.
func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{App: bot}, nil
}

func init() {
	err := app.RegisterModule("Username", NewModule)
	if err != nil {
		panic(err)
	}
}
