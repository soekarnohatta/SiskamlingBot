package admin

import (
	"SiskamlingBot/bot/core/telegram"
	misc "SiskamlingBot/bot/modules/misc"
	"SiskamlingBot/bot/utils"
	"encoding/json"
	"fmt"
)

func (m Module) getUser(ctx *telegram.TgContext) error {
	if len(ctx.Args()) < 1 {
		ctx.ReplyMessage("Pengguna tidak ditemukan!")
		return nil
	}

	usr, err := m.App.DB.User.GetUserById(utils.StrToInt64(ctx.Args()[0]))
	if usr != nil {
		infoUser := fmt.Sprintf(
			"<b>Info Pengguna</b>"+
				"\n<b>User ID</b>: <code>%v</code>"+
				"\n<b>Username</b>: <code>%v</code>"+
				"\n<b>First Name</b>: <code>%v</code>"+
				"\n<b>Last Name</b>: <code>%v</code>"+
				"\n<b>Is Banned</b>: <code>%v</code>",
			ctx.Args()[0],
			usr.UserName,
			usr.FirstName,
			usr.LastName,
			misc.Module{App: m.App}.IsBan(utils.StrToInt64(ctx.Args()[0])))

		ctx.ReplyMessage(infoUser)
		return nil
	}

	return err
}

func (m Module) getChat(ctx *telegram.TgContext) error {
	if len(ctx.Args()) < 1 {
		ctx.ReplyMessage("Obrolan tidak ditemukan!")
		return nil
	}

	cht, err := m.App.DB.Chat.GetChatById(utils.StrToInt64(ctx.Args()[0]))
	if cht != nil {
		infoChat := fmt.Sprintf(
			"<b>Info Obrolan</b>"+
				"\n<b>Chat ID</b>: <code>%v</code>"+
				"\n<b>Chat Name</b>: <code>%v</code>"+
				"\n<b>Chat Invitelink</b>: <code>%v</code>"+
				"\n<b>Chat Type</b>: <code>%v</code>",
			ctx.Args()[0],
			cht.ChatTitle,
			cht.ChatLink,
			cht.ChatType)

		ctx.ReplyMessage(infoChat)
		return nil
	}
	return err
}

func (Module) debug(ctx *telegram.TgContext) error {
	if ctx.Message.ReplyToMessage != nil {
		output, _ := json.MarshalIndent(ctx.Context.Update, "", "  ")
		ctx.ReplyMessage(fmt.Sprintf("<code>%s</code>", string(output)))
		return nil
	}

	output, _ := json.MarshalIndent(ctx.Context.Update, "", "  ")
	ctx.ReplyMessage(fmt.Sprintf("<code>%s</code>", string(output)))
	return nil
}
