package app

import (
	"SiskamlingBot/bot/core/telegram"
	"log"
)

type ModuleInfo struct {
	Name string
}

type Module interface {
	Info() ModuleInfo
	Commands() []telegram.Command
	Messages() []telegram.Message
	Callbacks() []telegram.Callback
}

type ModuleConstructor func(*MyApp) (Module, error)

var Modules = make(map[string]ModuleConstructor)

func RegisterModule(name string, constructor ModuleConstructor) {
	if _, ok := Modules[name]; ok {
		log.Fatalf("Attempted to register module under occupied name '%s'", name)
	}

	Modules[name] = constructor
	log.Printf("%s module has been loaded succesfully!\n", name)
}