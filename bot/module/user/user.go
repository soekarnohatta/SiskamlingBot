package user

import (
	"SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
)

// Module contains the state for an instance of this module.
type Module struct {
	Bot *core.TelegramBot
}

// Info returns basic information about this module.
func (m Module) Info() core.ModuleInfo {
	return core.ModuleInfo{
		Name: "User",
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
	}
}

func (m Module) Messages() []telegram.Message {
	return []telegram.Message{
		{
			Name:        "UsernameScanner",
			Description: "Detect user without username",
			Func:        m.usernameScan,
		},
		{
			Name:        "PictureScanner",
			Description: "Detect user without profile picture",
			Func:        m.pictureScan,
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
		{
			Name:        "PictureCallback",
			Description: "",
			Callback:    `picture\((.+?)\)`,
			Func:        m.pictureCallback,
		},
	}
}

// NewModule returns a new instance of this module.
func NewModule(bot *core.TelegramBot) (core.Module, error) {
	return &Module{
		Bot: bot,
	}, nil
}

func init() {
	core.RegisterModule("User", NewModule)
}
