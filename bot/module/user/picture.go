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

const picLog = `#PICTURE
<b>User Name:</b> %s
<b>User ID:</b> <code>%v</code>
<b>Chat Name:</b> %s
<b>Chat ID:</b> <code>%v</code>
<b>Link:</b> %s`

func (m Module) pictureScan(ctx *telegram.TgContext) {
	// WIP 16/05/2020
	/*
		member, err := ctx.Bot.GetChatMember(ctx.Message.Chat.Id, ctx.Message.From.Id)
		if err != nil {
			log.Println("failed to GetChatMember: " + err.Error())
			return
		}
	*/

	// WIP 16/05/2020
	if getStatus, _ := model.GetPictureByID(m.App.DB, context.TODO(), ctx.User.Id); /*member.CanSendMessages == false ||*/
	getStatus != nil &&
		getStatus.ChatID == ctx.Chat.Id &&
		getStatus.IsMuted {
		return
	}

	err := model.SavePicture(m.App.DB, context.TODO(), model.NewPicture(
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
	textToSend := fmt.Sprintf("⚠ Pengguna <b>%v</b> [<code>%v</code>] telah dibisukan karena belum memasang <b>Foto Profil!</b>",
		util.MentionHtml(int(ctx.Message.From.Id), ctx.Message.From.FirstName),
		ctx.Message.From.Id)
	ctx.SendMessageKeyboard(textToSend, 0, util.BuildKeyboardf(
		"./data/keyboard/picture.json",
		1,
		map[string]string{
			"1": strconv.Itoa(int(ctx.Message.From.Id)),
		}))

	txtToSend := fmt.Sprintf(picLog,
		util.MentionHtml(int(ctx.User.Id), ctx.User.FirstName),
		ctx.User.Id,
		ctx.Chat.Title,
		ctx.Chat.Id,
		util.CreateLinkHtml(util.CreateMessageLink(ctx.Chat, ctx.Message.MessageId), "Here"))

	ctx.SendMessage(txtToSend, m.App.Config.LogEvent)
}

func (m Module) pictureCallback(ctx *telegram.TgContext) {
	pattern, _ := regexp.Compile(`picture\((.+?)\)`)
	if !(pattern.FindStringSubmatch(ctx.Callback.Data)[1] == strconv.Itoa(int(ctx.Callback.From.Id))) {
		ctx.AnswerCallback("❌ ANDA BUKAN PENGGUNA YANG DIMAKSUD!", true)
		return
	}

	if p, err := ctx.Callback.From.GetProfilePhotos(ctx.Bot, nil); p != nil && p.TotalCount == 0 {
		if err != nil {
			log.Println("failed to get pictures: " + err.Error())
			return
		}

		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG FOTO PROFIL", true)
		return
	}

	// Delete user status if user has set username
	err := model.DeletePictureByID(m.App.DB, context.TODO(), ctx.Callback.From.Id)
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

	ctx.AnswerCallback("✅ Terimakasih telah memasang Foto Profil", true)
	ctx.DeleteMessage(0)
}
