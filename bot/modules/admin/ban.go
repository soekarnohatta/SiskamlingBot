package admin

import (
	"fmt"

	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
)

func (m Module) globalBan(ctx *telegram.TgContext) error {
	if !telegram.IsSudo(ctx.User.Id, m.App.Config.SudoUsers) {
		return nil
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("Tolong masukan id pengguna!", 0)
		return nil
	}

	var success int
	var failed int
	for idx, val := range ctx.Args() {
		if !utils.IsNumericOnly(val) {
			continue
		}

		var text = fmt.Sprintf("Starting global ban of <code>%v</code> ...", val)
		var getUser, _ = m.App.DB.User.GetUserById(utils.StrToInt64(val))

		if idx < 1 {
			ctx.SendMessage(text, 0)
		} else {
			ctx.EditMessage(text)
		}

		if getUser != nil {
			getUser.Gban = true
			var err = m.App.DB.User.SaveUser(getUser)
			if err != nil {
				text = fmt.Sprintf("Global ban of <code>%v</code> has failed due to: %s", val, err.Error())
				ctx.EditMessage(text)
				return err
			}

			text = fmt.Sprintf("Succeed! <code>%v</code> has been banned.", ctx.Args()[0])
			ctx.EditMessage(text)
			success++
			continue
		}

		failed++
		text = fmt.Sprintf("Global ban of <code>%v</code> has failed due to record not found", val)
		ctx.EditMessage(text)
		continue
	}

	var text = fmt.Sprintf("<code>%v</code> individuals have been banned, %v has failed.", success, failed)
	ctx.SendMessage(text, 0)
	return nil
}

func (m Module) removeGlobalBan(ctx *telegram.TgContext) error {
	if !telegram.IsSudo(ctx.User.Id, m.App.Config.SudoUsers) {
		return nil
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("Tolong masukan id pengguna!", 0)
		return nil
	}

	var text = fmt.Sprintf("Starting to remove global ban of <code>%v</code> ...", ctx.Args()[0])
	var getUser, err = m.App.DB.User.GetUserById(utils.StrToInt64(ctx.Args()[0]))

	ctx.SendMessage(text, 0)
	if getUser != nil && getUser.Gban {
		getUser.Gban = false
		var err = m.App.DB.User.SaveUser(getUser)
		if err != nil {
			return err
		}

		text = fmt.Sprintf("Succeed! <code>%v</code> has been unbanned.", ctx.Args()[0])
		ctx.EditMessage(text)
		return nil
	} else if getUser != nil && !getUser.Gban {
		text = fmt.Sprintf("<code>%v</code> has not been banned! anywhere", ctx.Args()[0])
		ctx.EditMessage(text)
		return nil
	}

	text = fmt.Sprintf("Global ban removal of <code>%v</code> has failed due to record not found", ctx.Args()[0])
	ctx.EditMessage(text)
	return err
}
