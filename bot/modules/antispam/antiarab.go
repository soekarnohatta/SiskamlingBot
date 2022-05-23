package user

import (
	"SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
	"fmt"
	"sync"
	"unicode"
)

func (m *Module) antiarab(ctx *telegram.TgContext) error {
	getPref, _ := m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if getPref != nil && !getPref.EnforceAntiArab {
		return telegram.ContinueOrder
	}

	getChat, _ := m.App.DB.Chat.GetChatById(ctx.Chat.Id)
	if getChat != nil && !getChat.ChatVIP {
		return telegram.ContinueOrder
	}

	if !isArab(ctx) {
		return telegram.ContinueOrder
	}

	toDelete := ctx.Message.MessageId
	text := fmt.Sprintf(
		"⚠ <b>%v</b> [<code>%v</code>] telah dihapus pesannya karena mengirim/menggunakan "+
			"karakter <b>Arabic</b>. Silahkan gunakan karakter lain.",
		telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName),
		ctx.User.Id,
	)

	banLog := fmt.Sprintf(
		"#ARAB"+
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

	wg.Add(3)
	go func() { defer wg.Done(); ctx.DeleteMessage(getPref.LastServiceMessageId) }()

	ctx.SendMessage(text, 0)
	getPref.LastServiceMessageId = ctx.Message.MessageId
	_ = m.App.DB.Pref.SavePreference(getPref)

	go func() { defer wg.Done(); ctx.DeleteMessage(toDelete) }()
	go func() { defer wg.Done(); ctx.SendMessageAsync(banLog, m.App.Config.LogEvent, nil) }()
	return telegram.EndOrder
}

func (m *Module) antiarabSetting(ctx *telegram.TgContext) error {
	if !core.IsUserAdmin(ctx) {
		ctx.SendMessage("Anda bukan admin!", 0)
		return nil
	}

	getChat, _ := m.App.DB.Chat.GetChatById(ctx.Chat.Id)
	if getChat != nil && !getChat.ChatVIP {
		ctx.SendMessage("Chat ini bukan VIP, silahkan beli VIP dahulu!", 0)
		return nil
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("Masukan argumen true/false", 0)
		return nil
	}

	prefs, err := m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if err != nil {
		ctx.SendMessage("Error pas ngambil data, coba lagi.", 0)
		return nil
	}

	extractArgs := utils.ExtractBool(ctx.Args()[0])
	prefs.EnforceAntiArab = extractArgs
	err = m.App.DB.Pref.SavePreference(prefs)
	if err != nil {
		ctx.SendMessage("Error pas masukin data, coba lagi.", 0)
		return err
	}

	ctx.SendMessage(fmt.Sprintf("Pengaturan deteksi karakter arab diatur ke <code>%v</code> ", extractArgs), 0)
	return nil
}

func isArab(ctx *telegram.TgContext) bool {
	return checkArab(ctx.User.FirstName) ||
		checkArab(ctx.User.LastName) ||
		checkArab(ctx.Message.Text) ||
		checkArab(ctx.Message.Caption)
}

func checkArab(text string) bool {
	for _, rangeTxt := range text {
		if unicode.Is(unicode.Arabic, rangeTxt) {
			return true
		}
		continue
	}
	return false
}
