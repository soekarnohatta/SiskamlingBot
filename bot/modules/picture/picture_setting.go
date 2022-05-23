package picture

import (
	"SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
	"fmt"
)

func (m *Module) pictureSetting(ctx *telegram.TgContext) error {
	if !core.IsUserAdmin(ctx) {
		ctx.SendMessage("Anda bukan admin!", 0)
		return nil
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("Masukan argumen true/false", 0)
		return nil
	}

	prefs, err := m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if err != nil {
		ctx.SendMessage("Error pas ngambil data, coba lagi.", 0)
		return err
	}

	var extractArgs = utils.ExtractBool(ctx.Args()[0])
	var txtToSend = fmt.Sprintf("Pengaturan pengawasan profil gambar diatur ke <code>%v</code>", extractArgs)
	prefs.EnforcePicture = extractArgs
	err = m.App.DB.Pref.SavePreference(prefs)
	if err != nil {
		ctx.SendMessage("Error pas masukin data, coba lagi.", 0)
		return err
	}

	ctx.SendMessage(txtToSend, 0)
	return nil
}
