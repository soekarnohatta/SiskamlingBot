package admin

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram/types"
)

// Module contains the state for an instance of this module.
type Module struct {
	App *app.MyApp
}

// Info returns basic information about this module.
func (*Module) Info() app.ModuleInfo {
	return app.ModuleInfo{Name: "Admin"}
}

// Commands returns a list of telegram provided by this module.
func (m *Module) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "Get User",
			Trigger:     "getuser",
			Description: "Get specific user info",
			Func:        m.getUser,
		},
		{
			Name:        "Get Chat",
			Trigger:     "getchat",
			Description: "Get specific chat info",
			Func:        m.getChat,
		},
		{
			Name:        "Debug",
			Trigger:     "dbg",
			Description: "Prints JSON dump of a message update",
			Func:        m.debug,
		},
		{
			Name:        "Gban",
			Trigger:     "gban",
			Description: "Ban user across chats",
			Func:        m.globalBan,
		},
		{
			Name:        "UnGban",
			Trigger:     "ungban",
			Description: "Unban user across chats",
			Func:        m.removeGlobalBan,
		},
	}
}

func (*Module) Messages() []types.Message {
	return nil
}

func (*Module) Callbacks() []types.Callback {
	return nil
}

// NewModule returns a new instance of this module.
func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{App: bot}, nil
}

func init() {
	app.RegisterModule("Admin", NewModule)
}
