package admin

import (
	"SiskamlingBot/bot/core/telegram"
	misc "SiskamlingBot/bot/modules/antispam"
	"SiskamlingBot/bot/utils"
	"encoding/json"
	"fmt"
)

func (m *Module) getUser(ctx *telegram.TgContext) error {
	if len(ctx.Args()) < 1 {
		ctx.ReplyMessage("Pengguna tidak ditemukan!")
		return nil
	}

	// no need to check error
	var usr, _ = m.App.DB.User.GetUserById(utils.StrToInt64(ctx.Args()[0]))
	if usr != nil {
		var module = misc.Module{App: m.App}
		var infoUser = fmt.Sprintf(
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
			module.IsBan(utils.StrToInt64(ctx.Args()[0])))

		ctx.ReplyMessage(infoUser)
		return nil
	}

	ctx.ReplyMessage("Pengguna tidak ditemukan!")
	return nil
}

func (m *Module) getChat(ctx *telegram.TgContext) error {
	if len(ctx.Args()) < 1 {
		ctx.ReplyMessage("Obrolan tidak ditemukan!")
		return nil
	}

	// no need to check error
	var cht, _ = m.App.DB.Chat.GetChatById(utils.StrToInt64(ctx.Args()[0]))
	if cht != nil {
		var infoChat = fmt.Sprintf(
			"<b>Info Obrolan</b>"+
				"\n<b>Chat ID</b>: <code>%v</code>"+
				"\n<b>Chat Name</b>: <code>%v</code>"+
				"\n<b>Chat Invitelink</b>: <code>%v</code>"+
				"\n<b>Chat Type</b>: <code>%v</code>"+
				"\n<b>Chat VIP</b>: <code>%v</code>",
			ctx.Args()[0],
			cht.ChatTitle,
			cht.ChatLink,
			cht.ChatType,
			cht.ChatVIP,
		)

		ctx.ReplyMessage(infoChat)
		return nil
	}

	ctx.ReplyMessage("Obrolan tidak ditemukan!")
	return nil
}

func (*Module) debug(ctx *telegram.TgContext) error {
	if ctx.Message.ReplyToMessage != nil {
		var output, _ = json.MarshalIndent(ctx.Context.Update, "", "  ")
		ctx.ReplyMessage(fmt.Sprintf("<code>%s</code>", string(output)))
		return nil
	}

	var output, _ = json.MarshalIndent(ctx.Context.Update, "", "  ")
	ctx.ReplyMessage(fmt.Sprintf("<code>%s</code>", string(output)))
	return nil
}

func (m *Module) addVip(ctx *telegram.TgContext) error {
	if !telegram.IsSudo(ctx.User.Id, m.App.Config.SudoUsers) {
		ctx.DeleteMessage(0)
		return nil
	}

	if len(ctx.Args()) < 2 {
		ctx.SendMessage("Argumen kurang!", 0)
		return nil
	}

	var chtId = utils.StrToInt64(ctx.Args()[0])
	var chat, err = m.App.DB.Chat.GetChatById(chtId)
	if err != nil {
		ctx.SendMessage("Error pas ngambil data/data tidak ditemukan, coba lagi.", 0)
		return nil
	}

	var extractArgs = utils.ExtractBool(ctx.Args()[1])
	var txtToSend = fmt.Sprintf("Pengaturan VIP <code>%v</code> diatur ke <code>%v</code> ", chtId, extractArgs)

	chat.ChatVIP = extractArgs
	err = m.App.DB.Chat.SaveChat(chat)
	if err != nil {
		ctx.SendMessage("Error pas masukin data, coba lagi.", 0)
		return err
	}

	ctx.SendMessage(txtToSend, 0)
	return nil
}
