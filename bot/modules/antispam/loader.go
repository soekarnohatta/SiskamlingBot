package user

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"

	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/core/telegram/types"
)

type Module struct {
	App *app.MyApp
}

func (*Module) Info() app.ModuleInfo {
	return app.ModuleInfo{Name: "Misc"}
}

func (m *Module) Commands() []types.Command {
	return []types.Command{
		{
			Name:    "ping",
			Trigger: "ping",
			Func:    m.ping,
		},
		{
			Name:    "about",
			Trigger: "about",
			Func:    m.about,
		},
		{
			Name:    "start",
			Trigger: "start",
			Func:    m.start,
		},
		{
			Name:    "info",
			Trigger: "info",
			Func:    m.info,
		},
		{
			Name:    "addbl",
			Trigger: "addbl",
			Func:    m.blacklistAdd,
		},
		{
			Name:    "delbl",
			Trigger: "delbl",
			Func:    m.blacklistRemove,
		},
		{
			Name:    "setantispam",
			Trigger: "setantispam",
			Func:    m.antispamSetting,
		},
		{
			Name:    "setantichinese",
			Trigger: "setantichinese",
			Func:    m.antichineseSetting,
		},
		{
			Name:    "setantiarab",
			Trigger: "setantiarab",
			Func:    m.antiarabSetting,
		},
	}
}

func (m *Module) Messages() []types.Message {
	return []types.Message{
		{
			Name:   "antispam",
			Filter: telegram.GroupFilter,
			Func:   m.antispam,
			Order:  0,
			Async:  false,
		},
		{
			Name:   "antiarab",
			Filter: telegram.GroupFilter,
			Func:   m.antiarab,
			Order:  0,
			Async:  true,
		},
		{
			Name:   "antichinese",
			Filter: telegram.GroupFilter,
			Func:   m.antichinese,
			Order:  0,
			Async:  true,
		},
		{
			Name:   "blacklist",
			Filter: message.All,
			Func:   m.blacklist,
			Order:  0,
			Async:  true,
		},
	}
}

func (m *Module) Callbacks() []types.Callback {
	return []types.Callback{
		{
			Name:     "HelpCallback",
			Callback: `help\((.+?)\)`,
			Func:     m.helpCallback,
		},
	}
}

func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{App: bot}, nil
}

func init() {
	err := app.RegisterModule("Misc", NewModule)
	if err != nil {
		panic(err)
	}
}
