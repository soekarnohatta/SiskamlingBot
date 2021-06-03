package core

import (
	"SiskamlingBot/bot/core/telegram"
	"fmt"
	"log"
	"strings"
)

func (b *MyApp) registerCommand(cmd telegram.Command) error {
	lName := strings.ToLower(cmd.Name)
	if _, ok := b.Commands[lName]; ok {
		return fmt.Errorf("register command '%s': name already used", cmd.Name)
	}
	b.Commands[lName] = cmd

	for _, alias := range cmd.Aliases {
		lAlias := strings.ToLower(alias)
		if _, ok := b.Commands[lAlias]; ok {
			return fmt.Errorf("register alias '%s' for command '%s': name already used", alias, cmd.Name)
		}
		b.Commands[lAlias] = cmd
	}

	return nil
}

func (b *MyApp) registerCommands(mod Module) error {
	for _, cmd := range mod.Commands() {
		if cmd.Func == nil {
			continue
		}

		err := b.registerCommand(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *MyApp) registerMessage(msg telegram.Message) error {
	lName := strings.ToLower(msg.Name)
	if _, ok := b.Messages[lName]; ok {
		return fmt.Errorf("register message '%s': name already used", msg.Name)
	}

	b.Messages[lName] = msg

	return nil
}

func (b *MyApp) registerMessages(mod Module) error {
	for _, cmd := range mod.Messages() {
		if cmd.Func == nil {
			continue
		}

		err := b.registerMessage(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *MyApp) registerCallback(cb telegram.Callback) error {
	lName := strings.ToLower(cb.Name)
	if _, ok := b.Messages[lName]; ok {
		return fmt.Errorf("register callback '%s': name already used", cb.Name)
	}

	b.Callbacks[lName] = cb
	return nil
}

func (b *MyApp) registerCallbacks(mod Module) error {
	for _, cmd := range mod.Callbacks() {
		if cmd.Func == nil {
			continue
		}

		err := b.registerCallback(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *MyApp) loadModule(name string, cstr ModuleConstructor) error {
	mod, err := cstr(b)
	if err != nil {
		return err
	}

	b.Modules[name] = mod

	err = b.registerCommands(mod)
	if err != nil {
		return err
	}

	err = b.registerMessages(mod)
	if err != nil {
		return err
	}

	err = b.registerCallbacks(mod)
	if err != nil {
		return err
	}

	return nil
}

func (b *MyApp) LoadModules() error {
	for name, cstr := range Modules {
		err := b.loadModule(name, cstr)
		if err != nil {
			return fmt.Errorf("load module '%s': %w", name, err)
		}
	}

	log.Println("Loaded All Modules")
	return nil
}
