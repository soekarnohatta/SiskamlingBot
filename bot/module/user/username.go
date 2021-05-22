package user

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/model"
	"SiskamlingBot/bot/util"
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

const unameLog = `#USERNAME
<b>User Name:</b> %s
<b>User ID:</b> <code>%v</code>
<b>Chat Name:</b> %s
<b>Chat ID:</b> <code>%v</code>
<b>Link:</b> %s`

func (m Module) usernameScan(ctx *telegram.TgContext) {
	// To avoid sending repeated message
	/*
		member, err := ctx.Bot.GetChatMember(ctx.Message.Chat.Id, ctx.Message.From.Id)
		if err != nil {
			log.Println("failed to GetChatMember: " + err.Error())
			return
		}
	*/

	// WIP 16/05/2020
	if getStatus, _ := model.GetUsernameByID(m.App.DB, context.TODO(), ctx.Message.From.Id);
	/*!member.CanSendMessages ||*/
	getStatus != nil &&
		getStatus.ChatID == ctx.Message.Chat.Id &&
		getStatus.IsMuted {
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
		util.MentionHtml(int(ctx.Message.From.Id), ctx.Message.From.FirstName),
		ctx.Message.From.Id)
	ctx.SendMessageKeyboard(textToSend, 0, util.BuildKeyboardf(
		"./data/keyboard/username.json",
		1,
		map[string]string{
			"1": strconv.Itoa(int(ctx.Message.From.Id)),
		}))

	txtToSend := fmt.Sprintf(unameLog,
		util.MentionHtml(int(ctx.User.Id), ctx.User.FirstName),
		ctx.User.Id,
		ctx.Chat.Title,
		ctx.Chat.Id,
		util.CreateLinkHtml(util.CreateMessageLink(ctx.Chat, ctx.Message.MessageId), "Here"))

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

	_, err := ctx.Bot.RestrictChatMember(ctx.Callback.Message.Chat.Id, ctx.Callback.From.Id, gotgbot.ChatPermissions{
		CanSendMessages:      true,
		CanSendMediaMessages: true,
		CanSendPolls:         true,
		CanSendOtherMessages: true,
	}, nil)
	if err != nil {
		log.Println("failed to unrestrict chatmember: " + err.Error())
		return
	}

	// Delete user status if user has set username
	err = model.DeleteUsernameByID(m.App.DB, context.TODO(), ctx.Callback.From.Id)
	if err != nil {
		log.Println("failed to save status to DB: " + err.Error())
		return
	}

	_, err = ctx.Bot.RestrictChatMember(ctx.Chat.Id, ctx.User.Id, gotgbot.ChatPermissions{
		CanSendMessages:      true,
		CanSendMediaMessages: true,
		CanSendPolls:         true,
		CanSendOtherMessages: true,
	}, nil)
	if err != nil {
		log.Println("failed to restrict chatmember: " + err.Error())
		return
	}

	ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
	ctx.DeleteMessage(0)
}
