package picture

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/core/telegram/types"
)

// Module contains the state for an instance of this module.
type Module struct {
	App *app.MyApp
}

// Info returns basic information about this module.
func (Module) Info() app.ModuleInfo {
	return app.ModuleInfo{Name: "Picture"}
}

// Commands returns a list of telegram provided by this module.
func (m Module) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "setpicture",
			Trigger:     "setpicture",
			Description: "ping the bot.",
			Func:        m.pictureSetting,
		},
	}
}

func (m Module) Messages() []types.Message {
	return []types.Message{
		{
			Name:        "PictureScanner",
			Description: "Detect user without profile picture",
			Filter:      telegram.ProfileAndGroupFilter(m.App.Bot),
			Func:        m.pictureScan,
			Order:       2,
			Async:       false,
		},
	}
}

func (m Module) Callbacks() []types.Callback {
	return []types.Callback{
		{
			Name:        "PictureCallback",
			Description: "",
			Callback:    `picture\((.+?)\)`,
			Func:        m.pictureCallback,
		},
	}
}

func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{App: bot}, nil
}

func init() {
	err := app.RegisterModule("Picture", NewModule)
	if err != nil {
		panic(err)
	}
}
