package app

import (
	"SiskamlingBot/bot/core/telegram/types"
	"fmt"
	"log"
	"strings"
)

func (b *MyApp) registerCommand(cmd types.Command) error {
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
			return fmt.Errorf("registerCallbacks: failed to register new command %s with error: %w", cmd.Name, err)
		}
	}

	return nil
}

func (b *MyApp) registerMessage(msg types.Message) error {
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
			return fmt.Errorf("registerMessages: failed to register new message %s with error: %w", cmd.Name, err)
		}
	}

	return nil
}

func (b *MyApp) registerCallback(cb types.Callback) error {
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
			return fmt.Errorf("registerCallbacks: failed to register new callback %s with error: %w", cmd.Name, err)
		}
	}

	return nil
}

func (b *MyApp) loadModule(name string, cstr ModuleConstructor) error {
	mod, err := cstr(b)
	if err != nil {
		return fmt.Errorf("loadModule: failed to construct new module with error: %w", err)
	}

	b.Modules[name] = mod

	err = b.registerCommands(mod)
	if err != nil {
		return fmt.Errorf("loadModule: failed to register new command with error: %w", err)
	}

	err = b.registerMessages(mod)
	if err != nil {
		return fmt.Errorf("loadModule: failed to register new message with error: %w", err)
	}

	err = b.registerCallbacks(mod)
	if err != nil {
		return fmt.Errorf("loadModule: failed to register new callback with error: %w", err)
	}

	return nil
}

func (b *MyApp) loadModules() error {
	for name, cstr := range Modules {
		err := b.loadModule(name, cstr)
		if err != nil {
			return fmt.Errorf("loadModules: failed to load module %s with error: %w", name, err)
		}
	}

	log.Println("Loaded All Modules!")
	return nil
}
