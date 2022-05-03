package app

import (
	"SiskamlingBot/bot/core/telegram/types"
	"fmt"
)

type ModuleInfo struct {
	Name string
}

type Module interface {
	Info() ModuleInfo
	Commands() []types.Command
	Messages() []types.Message
	Callbacks() []types.Callback
}

type ModuleConstructor func(*MyApp) (Module, error)

var Modules = make(map[string]ModuleConstructor)

func RegisterModule(name string, constructor ModuleConstructor) error {
	if _, ok := Modules[name]; ok {
		return fmt.Errorf("Attempted to register module under occupied name '%s'", name)
	}

	Modules[name] = constructor
	return nil
}
