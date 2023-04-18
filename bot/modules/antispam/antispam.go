package user

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
	"fmt"
	"sync"
)

func (m *Module) antispam(ctx *telegram.TgContext) error {
	var getPref, _ = m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if getPref != nil && !getPref.EnforceAntispam {
		return telegram.ContinueOrder
	}

	var user = ctx.User
	if !m.IsBan(user.Id) {
		return telegram.ContinueOrder
	}

	var banChan = make(chan bool, 1)
	go func() { banChan <- ctx.BanChatMember(0, 0) }()

	var dataMap = map[string]string{
		"1": telegram.MentionHtml(user.Id, user.FirstName),
		"2": utils.Int64ToStr(user.Id),
	}

	var toDeleteServiceMessage = getPref.LastServiceMessageId
	var toDeleteAndSave = ctx.Message.MessageId
	var text, keyb = telegram.CreateMenuf("./data/menu/spam.json", 1, dataMap)
	var banLog = fmt.Sprintf(
		"#BAN"+
			"\n<b>User Name:</b> %s"+
			"\n<b>User ID:</b> <code>%v</code>"+
			"\n<b>Chat Name:</b> %s"+
			"\n<b>Chat ID:</b> <code>%v</code>"+
			"\n<b>Link:</b> %s",
		telegram.MentionHtml(user.Id, user.FirstName),
		user.Id,
		ctx.Chat.Title,
		ctx.Chat.Id,
		telegram.CreateLinkHtml(telegram.CreateMessageLink(ctx.Chat, toDeleteAndSave), "Here"),
	)

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go func() { defer wg.Done(); ctx.DeleteMessage(toDeleteServiceMessage) }()

	if !<-banChan {
		text += "\n\n🚫 <b>Tetapi saya tidak bisa mengeluarkannya, mohon periksa kembali perizinan saya!</b>"
		ctx.SendMessage(text, 0)
		getPref.LastServiceMessageId = ctx.Message.MessageId
		var err = m.App.DB.Pref.SavePreference(getPref)
		if err != nil {
			return err
		}

		return telegram.EndOrder
	}

	wg.Add(3)
	go func() {
		ctx.SendMessageKeyboard(text, 0, keyb)
		getPref.LastServiceMessageId = ctx.Message.MessageId
		_ = m.App.DB.Pref.SavePreference(getPref)
		wg.Done()
	}()

	go func() { ctx.DeleteMessage(toDeleteAndSave); wg.Done() }()
	go func() { ctx.SendMessageKeyboardAsync(banLog, m.App.Config.LogEvent, nil); wg.Done() }()
	return telegram.EndOrder
}

func (m *Module) antispamSetting(ctx *telegram.TgContext) error {
	if !utils.IsUserAdmin(ctx.Bot, ctx.Chat.Id, ctx.User.Id) {
		ctx.SendMessage("Anda bukan admin!", 0)
		return nil
	}

	if len(ctx.Args()) < 1 {
		ctx.SendMessage("Masukan argumen true/false", 0)
		return nil
	}

	var prefs, err = m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if err != nil {
		ctx.SendMessage("Error pas ngambil data, coba lagi.", 0)
		return err
	}

	var extractArgs = utils.ExtractBool(ctx.Args()[0])
	prefs.EnforceAntispam = extractArgs
	err = m.App.DB.Pref.SavePreference(prefs)
	if err != nil {
		ctx.SendMessage("Error pas masukin data, coba lagi.", 0)
		return err
	}

	var toSend = fmt.Sprintf("Pengaturan antispam diatur ke <code>%v</code> ", extractArgs)
	ctx.SendMessage(toSend, 0)
	return nil
}
