package user

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
	"fmt"
	"sync"
)

func (m Module) antispam(ctx *telegram.TgContext) error {
	getPref, _ := m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if getPref != nil && !getPref.EnforceAntispam {
		return telegram.ContinueOrder
	}

	user := ctx.User
	if !m.IsBan(user.Id) {
		return telegram.ContinueOrder
	}

	dataMap := map[string]string{
		"1": telegram.MentionHtml(user.Id, user.FirstName),
		"2": utils.Int64ToStr(user.Id),
	}

	text, keyb := telegram.CreateMenuf("./data/menu/spam.json", 1, dataMap)
	banLog := fmt.Sprintf(
		"#BAN"+
			"\n<b>User Name:</b> %s"+
			"\n<b>User ID:</b> <code>%v</code>"+
			"\n<b>Chat Name:</b> %s"+
			"\n<b>Chat ID:</b> <code>%v</code>"+
			"\n<b>Link:</b> %s",
		telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName),
		ctx.User.Id,
		ctx.Chat.Title,
		ctx.Chat.Id,
		telegram.CreateLinkHtml(telegram.CreateMessageLink(ctx.Chat, ctx.Message.MessageId), "Here"),
	)

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go func() { defer wg.Done(); ctx.DeleteMessage(getPref.LastServiceMessageId) }()

	if !ctx.BanChatMember(0, 0) {
		text += "\n\n🚫 <b>Tetapi saya tidak bisa mengeluarkannya, mohon periksa kembali perizinan saya!</b>"
		ctx.SendMessage(text, 0)
		getPref.LastServiceMessageId = ctx.Message.MessageId
		err := m.App.DB.Pref.SavePreference(getPref)
		if err != nil {
			return err
		}

		return telegram.EndOrder
	}

	ctx.DeleteMessage(0)
	ctx.SendMessage(banLog, m.App.Config.LogEvent)
	ctx.SendMessageKeyboard(text, 0, keyb)
	getPref.LastServiceMessageId = ctx.Message.MessageId
	_ = m.App.DB.Pref.SavePreference(getPref)
	return telegram.EndOrder
}
