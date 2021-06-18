package admin

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/model"
	"SiskamlingBot/bot/util"
	"context"
	"fmt"
)

const (
	infoUser= `<b>Info Pengguna</b>
<b>User ID</b>: <code>%v</code>
<b>Username</b>: <code>%v</code>
<b>First Name</b>: <code>%v</code>
<b>Last Name</b>: <code>%v</code>`

	infoChat = `<b>Info Obrolan</b>
<b>Chat ID</b>: <code>%v</code>
<b>Chat Name</b>: <code>%v</code>
<b>Chat Invitelink</b>: <code>%v</code>
<b>Chat Type</b>: <code>%v</code>`
)

func (m Module) getUser(ctx *telegram.TgContext) {
	if len(ctx.Args()) > 0 {
		usr, err := model.GetUserByID(m.App.DB, context.TODO(), util.StrToInt(ctx.Args()[0]))
		if usr != nil && err == nil {
			formattedText := fmt.Sprintf(infoUser, ctx.Args()[0], usr.UserName, usr.FirstName, usr.LastName)
			ctx.ReplyMessage(formattedText)
			return
		}
	}

	ctx.ReplyMessage("Pengguna tidak ditemukan!")
}

func (m Module) getChat(ctx *telegram.TgContext) {
	if len(ctx.Args()) > 0 {
		cht, err := model.GetChatByID(m.App.DB, context.TODO(), util.StrToInt(ctx.Args()[0]))
		if cht != nil && err == nil {
			formattedText := fmt.Sprintf(infoChat, ctx.Args()[0], cht.ChatTitle, cht.ChatLink, cht.ChatType)
			ctx.ReplyMessage(formattedText)
			return
		}
	}

	ctx.ReplyMessage("Obrolan tidak ditemukan!")
}
