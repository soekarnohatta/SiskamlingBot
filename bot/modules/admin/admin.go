package admin

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/models"
	"SiskamlingBot/bot/utils"
	"encoding/json"
	"fmt"
)

const (
	infoUser = `<b>Info Pengguna</b>
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
		usr, err := models.GetUserByID(m.App.DB, utils.StrToInt(ctx.Args()[0]))
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
		cht, err := models.GetChatByID(m.App.DB, utils.StrToInt(ctx.Args()[0]))
		if cht != nil && err == nil {
			formattedText := fmt.Sprintf(infoChat, ctx.Args()[0], cht.ChatTitle, cht.ChatLink, cht.ChatType)
			ctx.ReplyMessage(formattedText)
			return
		}
	}

	ctx.ReplyMessage("Obrolan tidak ditemukan!")
}

func (m Module) debug(ctx *telegram.TgContext) {
	if ctx.Message.ReplyToMessage != nil {
		output, _ := json.MarshalIndent(ctx.Message.ReplyToMessage, "", "  ")
		ctx.ReplyMessage(string(output))
		return
	}

	output, _ := json.MarshalIndent(ctx.Message, "", "  ")
	ctx.ReplyMessage(string(output))
}
