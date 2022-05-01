package admin

import (
	"SiskamlingBot/bot/core/telegram"
	misc "SiskamlingBot/bot/modules/misc"
	"SiskamlingBot/bot/utils"
	"encoding/json"
	"fmt"
)

const (
	infoUser = `<b>Info Pengguna</b>
  <b>User ID</b>: <code>%v</code>
  <b>Username</b>: <code>%v</code>
  <b>First Name</b>: <code>%v</code>
  <b>Last Name</b>: <code>%v</code>
  <b>Is Banned</b>: <code>%v</code>`

	infoChat = `<b>Info Obrolan</b>
  <b>Chat ID</b>: <code>%v</code>
  <b>Chat Name</b>: <code>%v</code>
  <b>Chat Invitelink</b>: <code>%v</code>
  <b>Chat Type</b>: <code>%v</code>`
)

func (m Module) getUser(ctx *telegram.TgContext) error {
	if len(ctx.Args()) < 1 {
		ctx.ReplyMessage("Pengguna tidak ditemukan!")
		return nil
	}

	usr, err := m.App.DB.User.GetUserById(utils.StrToInt64(ctx.Args()[0]))
	if usr != nil {
		formattedText := fmt.Sprintf(infoUser, ctx.Args()[0], usr.UserName, usr.FirstName, usr.LastName, misc.Module{App: m.App}.IsBan(utils.StrToInt64(ctx.Args()[0])))
		ctx.ReplyMessage(formattedText)
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
		formattedText := fmt.Sprintf(infoChat, ctx.Args()[0], cht.ChatTitle, cht.ChatLink, cht.ChatType)
		ctx.ReplyMessage(formattedText)
		return nil
	}
	return err
}

func (Module) debug(ctx *telegram.TgContext) error {
	if ctx.Message.ReplyToMessage != nil {
		output, _ := json.MarshalIndent(ctx.Message.ReplyToMessage, "", "  ")
		ctx.ReplyMessage(string(output))
		return nil
	}

	output, _ := json.MarshalIndent(ctx.Message, "", "  ")
	ctx.ReplyMessage(string(output))
	return nil
}
