package username

import (
	"SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/models"
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const (
	unameLog = `#USERNAME
<b>User Name:</b> %s
<b>User ID:</b> <code>%v</code>
<b>Chat Name:</b> %s
<b>Chat ID:</b> <code>%v</code>
<b>Link:</b> %s`

	unameMsg = "⚠ <b>%v</b> [<code>%v</code>] telah dibisukan karena belum memasang <b>Username!</b>"
)

func (m Module) usernameScan(ctx *telegram.TgContext) {
	if core.IsUserBotRestricted(ctx, m.App) {
		return
	}

	newUsername := models.NewUsername(ctx.User.Id, ctx.User.Id, true)
	err := models.SaveUsername(m.App.DB, context.TODO(), newUsername)
	if err != nil {
		log.Println("failed to save status to DB: " + err.Error())
		return
	}

	if !ctx.RestrictMember(0, 0) {
		return
	}

	ctx.DeleteMessage(0)
	textToSend := fmt.Sprintf(unameMsg, telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName), ctx.User.Id)
	ctx.SendMessageKeyboard(textToSend, 0, telegram.BuildKeyboardf("./data/keyboard/username.json", 1, map[string]string{"1": strconv.Itoa(int(ctx.User.Id))}))

	textToSend = fmt.Sprintf(unameLog,
		telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName),
		ctx.User.Id,
		ctx.Chat.Title,
		ctx.Chat.Id,
		telegram.CreateLinkHtml(telegram.CreateMessageLink(ctx.Chat, ctx.Message.MessageId), "Here"))

	ctx.SendMessage(textToSend, m.App.Config.LogEvent)
}

func (m Module) usernameCallback(ctx *telegram.TgContext) {
	pattern, _ := regexp.Compile(`username\((.+?)\)`)
	if !(pattern.FindStringSubmatch(ctx.Callback.Data)[1] == strconv.Itoa(int(ctx.Callback.From.Id))) {
		getUsername, _ := models.GetUsernameByID(m.App.DB, context.TODO(), ctx.Callback.From.Id)
		if getUsername != nil && getUsername.ChatID == ctx.Callback.Message.Chat.Id {
			if ctx.User.Username == "" {
				ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
				return
			}
		
			err := models.DeleteUsernameByID(m.App.DB, context.TODO(), ctx.Callback.From.Id)
			if err != nil {
				log.Println("failed to save status to DB: " + err.Error())
			}
		
			ctx.UnRestrictMember(0)
			ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
			//ctx.DeleteMessage(0)
			return 
		}

		ctx.AnswerCallback("❌ ANDA BUKAN PENGGUNA YANG DIMAKSUD!", true)
		return
	}

	if ctx.User.Username == "" {
		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
		return
	}

	err := models.DeleteUsernameByID(m.App.DB, context.TODO(), ctx.Callback.From.Id)
	if err != nil {
		log.Println("failed to save status to DB: " + err.Error())
	}

	ctx.UnRestrictMember(0)
	ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
	ctx.DeleteMessage(0)
}
