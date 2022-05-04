package user

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
	"fmt"
	"sync"
)

const banLog = `#BAN
<b>User Name:</b> %s
<b>User ID:</b> <code>%v</code>
<b>Chat Name:</b> %s
<b>Chat ID:</b> <code>%v</code>
<b>Link:</b> %s`

func (m Module) antispam(ctx *telegram.TgContext) error {
	getPref, err := m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if getPref != nil && !getPref.EnforceAntispam {
		return telegram.ContinueOrder
	} else if err != nil {
		return err
	}

	user := ctx.User
	if !m.IsBan(user.Id) {
		return telegram.ContinueOrder
	}

	dataMap := map[string]string{"1": telegram.MentionHtml(user.Id, user.FirstName), "2": utils.IntToStr(int(user.Id))}
	text, keyb := telegram.CreateMenuf("./data/menu/spam.json", 1, dataMap)

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func() { defer wg.Done(); ctx.DeleteMessage(getPref.LastServiceMessageId) }()
	if _, err := ctx.Bot.BanChatMember(ctx.Message.Chat.Id, user.Id, nil); err != nil {
		text += "\n\n🚫 <b>Tetapi saya tidak bisa mengeluarkannya, mohon periksa kembali perizinan saya!</b>"
		ctx.SendMessage(text, 0)
		getPref.LastServiceMessageId = ctx.Message.MessageId
		err := m.App.DB.Pref.SavePreference(getPref)
		if err != nil {
			return err
		}
		return telegram.EndOrder
	}

	wg.Add(3)
	go func() {
		defer wg.Done()
		ctx.SendMessageKeyboard(text, 0, keyb)
		getPref.LastServiceMessageId = ctx.Message.MessageId
		_ = m.App.DB.Pref.SavePreference(getPref)
	}()
	go func() { defer wg.Done(); ctx.DeleteMessage(0) }()
	go func() {
		defer wg.Done()
		textToSend := fmt.Sprintf(banLog,
			telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName),
			ctx.User.Id,
			ctx.Chat.Title,
			ctx.Chat.Id,
			telegram.CreateLinkHtml(telegram.CreateMessageLink(ctx.Chat, ctx.Message.MessageId), "Here"))
		ctx.SendMessage(textToSend, m.App.Config.LogEvent)
	}()
	return telegram.EndOrder
}
