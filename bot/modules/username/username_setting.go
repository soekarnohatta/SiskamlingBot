package username

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
	"fmt"
)

func (m *Module) usernameSetting(ctx *telegram.TgContext) error {
	if !utils.IsUserAdmin(ctx.Bot, ctx.Chat.Id, ctx.User.Id) {
		ctx.SendMessage("Anda bukan admin!", 0)
		return nil
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("Masukan argumen true/false", 0)
		return nil
	}

	var prefs, err = m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if err != nil {
		ctx.SendMessage("Error pas ngambil data, coba lagi.", 0)
		return err
	}

	var extractArgs = utils.ExtractBool(ctx.Args()[0])
	var txtToSend = fmt.Sprintf("Pengaturan pengawasan username diatur ke <code>%v</code> ", extractArgs)
	prefs.EnforceUsername = extractArgs
	err = m.App.DB.Pref.SavePreference(prefs)
	if err != nil {
		ctx.SendMessage("Error pas masukin data, coba lagi.", 0)
		return err
	}

	ctx.SendMessage(txtToSend, 0)
	return nil
}
