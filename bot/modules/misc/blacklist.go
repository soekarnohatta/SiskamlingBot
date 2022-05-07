package user

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/models"
	"fmt"
	"strings"
	"sync"
)

func (m Module) blacklist(ctx *telegram.TgContext) error {
	var wg sync.WaitGroup
	defer wg.Wait()

	var getBlacklist, err = m.App.DB.Blacklist.GetAllBlacklist()
	if err != nil {
		return err
	}

	for _, val := range getBlacklist {
		wg.Add(1)
		go func(compare string) {
			defer wg.Done()
			var text = strings.ToLower(ctx.Message.Text)
			var comp = strings.ToLower(compare)
			if strings.Contains(text, comp) {
				ctx.DeleteMessage(0)
			}
		}(val.BlacklistTrigger)
	}

	return telegram.ContinueOrder
}

func (m Module) blacklistAdd(ctx *telegram.TgContext) error {
	if !telegram.IsSudo(ctx.User.Id, m.App.Config.SudoUsers) {
		return nil
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("Masukan kata yang mau di blacklist", 0)
		return nil
	}

	var newBlacklist = &models.Blacklist{BlacklistTrigger: ctx.Args()[0]}
	var err = m.App.DB.Blacklist.SaveBlacklist(newBlacklist)
	if err != nil {
		return err
	}

	ctx.SendMessage(fmt.Sprintf("Trigger <code>%v</code> berhasil dimasukkan", ctx.Args()[0]), 0)
	return nil
}

func (m Module) blacklistRemove(ctx *telegram.TgContext) error {
	if !telegram.IsSudo(ctx.User.Id, m.App.Config.SudoUsers) {
		return nil
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("Masukan kata yang mau di hapus dari blacklist", 0)
		return nil
	}

	var err = m.App.DB.Blacklist.DeleteBlacklistByTrigger(ctx.Args()[0])
	if err != nil {
		return err
	}

	ctx.SendMessage(fmt.Sprintf("Trigger <code>%v</code> berhasil diapus", ctx.Args()[0]), 0)
	return nil
}
