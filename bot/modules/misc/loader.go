package user

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram/types"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
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
			Name:        "ping",
			Trigger:     "ping",
			Description: "ping the bot.",
			Func:        m.ping,
		},
		{
			Name:        "about",
			Trigger:     "about",
			Description: "about the bot.",
			Func:        m.about,
		},
		{
			Name:        "start",
			Trigger:     "start",
			Description: "start the bot.",
			Func:        m.start,
		},
		{
			Name:        "info",
			Trigger:     "info",
			Description: "info the bot.",
			Func:        m.info,
		},
		{
			Name:        "addbl",
			Trigger:     "addbl",
			Description: "info the bot.",
			Func:        m.blacklistAdd,
		},
		{
			Name:        "delbl",
			Trigger:     "delbl",
			Description: "info the bot.",
			Func:        m.blacklistRemove,
		},
		{
			Name:        "setantispam",
			Trigger:     "setantispam",
			Description: "info the bot.",
			Func:        m.antispamSetting,
		},
	}
}

func (m *Module) Messages() []types.Message {
	return []types.Message{
		{
			Name:        "antispam",
			Description: "Detect user without username",
			Filter:      message.All,
			Func:        m.antispam,
			Order:       1,
			Async:       false,
		},
		{
			Name:        "blacklist",
			Description: "Detect blacklisted trigger",
			Filter:      message.All,
			Func:        m.blacklist,
			Order:       0,
			Async:       true,
		},
	}
}

func (m *Module) Callbacks() []types.Callback {
	return []types.Callback{
		{
			Name:        "HelpCallback",
			Description: "",
			Callback:    `help\((.+?)\)`,
			Func:        m.helpCallback,
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
