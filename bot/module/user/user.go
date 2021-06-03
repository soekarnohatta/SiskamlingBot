package user

import (
	app "SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
)

// Module contains the state for an instance of this module.
type Module struct {
	App *app.MyApp
}

// Info returns basic information about this module.
func (m Module) Info() app.ModuleInfo {
	return app.ModuleInfo{
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
	return []telegram.Message{
		{
			Name:        "UsernameScanner",
			Description: "Detect user without username",
			Filter:      telegram.UsernameAndGroupFilter,
			Func:        m.usernameScan,
		},
		{
			Name:        "PictureScanner",
			Description: "Detect user without profile picture",
			Filter:      telegram.ProfileAndGroupFilter(m.App.Bot),
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
func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{
		App: bot,
	}, nil
}

func init() {
	app.RegisterModule("User", NewModule)
}
