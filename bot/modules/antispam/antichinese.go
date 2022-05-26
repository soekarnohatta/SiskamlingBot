package user

import (
	"fmt"
	"sync"
	"unicode"

	"SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
)

func (m *Module) antichinese(ctx *telegram.TgContext) error {
	var getPref, _ = m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if getPref != nil && !getPref.EnforceAntiChinese {
		return telegram.ContinueOrder
	}

	var getChat, _ = m.App.DB.Chat.GetChatById(ctx.Chat.Id)
	if getChat != nil && !getChat.ChatVIP {
		return telegram.ContinueOrder
	}

	if !isChinese(ctx) {
		return telegram.ContinueOrder
	}

	var toDeleteServiceMessage = getPref.LastServiceMessageId
	var toDeleteAndSave = ctx.Message.MessageId
	var text = fmt.Sprintf(
		"⚠ <b>%v</b> [<code>%v</code>] telah dihapus pesannya karena mengirim/menggunakan "+
			"karakter <b>Chinese</b>. Silahkan gunakan karakter lain.",
		telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName),
		ctx.User.Id,
	)

	var banLog = fmt.Sprintf(
		"#CHINESE"+
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
	wg.Add(4)

	go func() {
		ctx.SendMessage(text, 0)
		getPref.LastServiceMessageId = ctx.Message.MessageId
		_ = m.App.DB.Pref.SavePreference(getPref)
		wg.Done()
	}()

	go func() { ctx.DeleteMessage(toDeleteServiceMessage); wg.Done() }()
	go func() { ctx.DeleteMessage(toDeleteAndSave); wg.Done() }()
	go func() { ctx.SendMessageAsync(banLog, m.App.Config.LogEvent, nil); wg.Done() }()
	return telegram.EndOrder
}

func (m *Module) antichineseSetting(ctx *telegram.TgContext) error {
	if !core.IsUserAdmin(ctx) {
		ctx.SendMessage("Anda bukan admin!", 0)
		return nil
	}

	var getChat, _ = m.App.DB.Chat.GetChatById(ctx.Chat.Id)
	if getChat != nil && !getChat.ChatVIP {
		ctx.SendMessage("Chat ini bukan VIP, silahkan beli VIP dahulu!", 0)
		return nil
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("Masukan argumen true/false", 0)
		return nil
	}

	var prefs, err = m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if err != nil {
		ctx.SendMessage("Error pas ngambil data, coba lagi.", 0)
		return nil
	}

	var extractArgs = utils.ExtractBool(ctx.Args()[0])
	prefs.EnforceAntiChinese = extractArgs
	err = m.App.DB.Pref.SavePreference(prefs)
	if err != nil {
		ctx.SendMessage("Error pas masukin data, coba lagi.", 0)
		return err
	}

	ctx.SendMessage(fmt.Sprintf("Pengaturan deteksi karakter chinese diatur ke <code>%v</code> ", extractArgs), 0)
	return nil
}

func isChinese(ctx *telegram.TgContext) bool {
	return checkChinese(ctx.User.FirstName) ||
		checkChinese(ctx.User.LastName) ||
		checkChinese(ctx.Message.Text) ||
		checkChinese(ctx.Message.Caption)
}

func checkChinese(text string) bool {
	for _, rangeTxt := range text {
		if unicode.Is(unicode.Han, rangeTxt) {
			return true
		}
		continue
	}
	return false
}
