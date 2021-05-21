package admin

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/model"
	"SiskamlingBot/bot/util"
	"context"
	"fmt"
)

func (m Module) getUser(ctx *telegram.TgContext) {
	if len(ctx.Args()) > 0 {
		usr, _ := model.GetUserByID(m.App.DB, context.TODO(), util.StrToInt(ctx.Args()[0]))
		if usr != nil {
			text := `<b>Info Pengguna</b>
User ID: <code>%v</code>
Username: <code>%v</code>
First Name: <code>%v</code>
Last Name: <code>%v</code>
			`
			formattedText := fmt.Sprintf(text, ctx.Args()[0], usr.UserName, usr.FirstName, usr.LastName)
			ctx.ReplyMessage(formattedText)
			return
		}
	}

	ctx.ReplyMessage("<code>Pengguna tidak ditemukan!</code>")
}
