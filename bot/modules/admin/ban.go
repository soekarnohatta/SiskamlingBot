package admin

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/models"
	"SiskamlingBot/bot/utils"
	"fmt"
)

func (m Module) globalBan(ctx *telegram.TgContext) {
	if !telegram.IsSudo(ctx.User.Id, m.App.Config.SudoUsers) {
		ctx.SendMessage("Anda bukan Sudo!", 0)
		return
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("masukin id nya dong", 0)
		return
	}

	text := fmt.Sprintf("Starting global ban of <code>%v</code> ...", ctx.Args()[0])
	ctx.SendMessage(text, 0)

	getUser := models.GetUserByID(m.App.DB, utils.StrToInt64(ctx.Args()[0]))
	if getUser != nil {
		getUser.Gban = true
		models.SaveUser(m.App.DB, getUser)

		text = fmt.Sprintf("Succeed! <code>%v</code> has been banned.", ctx.Args()[0])
		ctx.EditMessage(text)
		return
	}

	text = fmt.Sprintf("Global ban of <code>%v</code> has failed due to record not found", ctx.Args()[0])
	ctx.EditMessage(text)
}

func (m Module) removeGlobalBan(ctx *telegram.TgContext) {
	if !telegram.IsSudo(ctx.User.Id, m.App.Config.SudoUsers) {
		ctx.SendMessage("Anda bukan Sudo!", 0)
		return
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("masukin id nya dong", 0)
		return
	}

	text := fmt.Sprintf("Starting to remove global ban of <code>%v</code> ...", ctx.Args()[0])
	ctx.SendMessage(text, 0)

	getUser := models.GetUserByID(m.App.DB, utils.StrToInt64(ctx.Args()[0]))
	if getUser != nil && getUser.Gban {
		getUser.Gban = false
		models.SaveUser(m.App.DB, getUser)

		text = fmt.Sprintf("Succeed! <code>%v</code> has been unbanned.", ctx.Args()[0])
		ctx.EditMessage(text)
		return
	} else if !getUser.Gban {
		text = fmt.Sprintf("<code>%v</code> has not been banned!", ctx.Args()[0])
		ctx.EditMessage(text)
		return
	}

	text = fmt.Sprintf("Global ban removal of <code>%v</code> has failed due to record not found", ctx.Args()[0])
	ctx.EditMessage(text)
}
