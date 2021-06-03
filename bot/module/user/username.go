package user

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/helper"
	"SiskamlingBot/bot/model"
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const unameLog = `#USERNAME
<b>User Name:</b> %s
<b>User ID:</b> <code>%v</code>
<b>Chat Name:</b> %s
<b>Chat ID:</b> <code>%v</code>
<b>Link:</b> %s`

func (m Module) usernameScan(ctx *telegram.TgContext) {
	if helper.IsUserBotRestricted(ctx, m.App) {
		return
	}

	err := model.SaveUsername(m.App.DB, context.TODO(), model.NewUsername(
		ctx.Message.From.Id,
		ctx.Message.Chat.Id,
		true,
	))
	if err != nil {
		log.Println("failed to save status to DB: " + err.Error())
		return
	}

	ctx.RestrictMember(0, 0)
	ctx.DeleteMessage(0)
	textToSend := fmt.Sprintf("⚠ Pengguna <b>%v</b> [<code>%v</code>] telah dibisukan karena belum memasang <b>Username!</b>",
		telegram.MentionHtml(int(ctx.Message.From.Id), ctx.Message.From.FirstName),
		ctx.Message.From.Id)
	ctx.SendMessageKeyboard(textToSend, 0, telegram.BuildKeyboardf(
		"./data/keyboard/username.json",
		1,
		map[string]string{
			"1": strconv.Itoa(int(ctx.Message.From.Id)),
		}))

	txtToSend := fmt.Sprintf(unameLog,
		telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName),
		ctx.User.Id,
		ctx.Chat.Title,
		ctx.Chat.Id,
		telegram.CreateLinkHtml(telegram.CreateMessageLink(ctx.Chat, ctx.Message.MessageId), "Here"))

	ctx.SendMessage(txtToSend, m.App.Config.LogEvent)
}

func (m Module) usernameCallback(ctx *telegram.TgContext) {
	pattern, _ := regexp.Compile(`username\((.+?)\)`)
	if !(pattern.FindStringSubmatch(ctx.Callback.Data)[1] == strconv.Itoa(int(ctx.Callback.From.Id))) {
		ctx.AnswerCallback("❌ ANDA BUKAN PENGGUNA YANG DIMAKSUD!", true)
		return
	}

	if ctx.User.Username == "" {
		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
		return
	}

	err := model.DeleteUsernameByID(m.App.DB, context.TODO(), ctx.Callback.From.Id)
	if err != nil {
		log.Println("failed to save status to DB: " + err.Error())
	}

	ctx.UnRestrictMember(0)
	ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
	ctx.DeleteMessage(0)
}
