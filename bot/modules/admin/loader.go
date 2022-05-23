package admin

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram/types"
)

type Module struct {
	App *app.MyApp
}

func (*Module) Info() app.ModuleInfo {
	return app.ModuleInfo{Name: "Admin"}
}

func (m *Module) Commands() []types.Command {
	return []types.Command{
		{
			Name:    "Get User",
			Trigger: "getuser",
			Func:    m.getUser,
		},
		{
			Name:    "Get Chat",
			Trigger: "getchat",
			Func:    m.getChat,
		},
		{
			Name:    "Debug",
			Trigger: "dbg",
			Func:    m.debug,
		},
		{
			Name:    "Gban",
			Trigger: "gban",
			Func:    m.globalBan,
		},
		{
			Name:    "UnGban",
			Trigger: "ungban",
			Func:    m.removeGlobalBan,
		},
		{
			Name:    "addvip",
			Trigger: "addvip",
			Func:    m.addVip,
		},
	}
}

func (*Module) Messages() []types.Message {
	return nil
}

func (*Module) Callbacks() []types.Callback {
	return nil
}

func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{App: bot}, nil
}

func init() {
	err := app.RegisterModule("Admin", NewModule)
	if err != nil {
		panic(err)
	}
}
