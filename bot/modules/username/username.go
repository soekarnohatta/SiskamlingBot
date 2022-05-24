package username

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
	"fmt"
	"regexp"
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (m *Module) usernameScan(ctx *telegram.TgContext) error {
	getPref, _ := m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if getPref != nil && !getPref.EnforceUsername {
		return telegram.ContinueOrder
	}

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(2)

	rstrChan := make(chan bool, 1)
	untilDate := utils.ExtractTime("5m")
	toDeleteServiceMessage := getPref.LastServiceMessageId
	toDeleteAndSave := ctx.Message.MessageId

	go func() { rstrChan <- ctx.RestrictMember(0, 0, untilDate) }()
	go func() { defer wg.Done(); ctx.DeleteMessage(toDeleteServiceMessage) }()

	dataButton := map[string]string{
		"1": utils.Int64ToStr(ctx.User.Id),
		"2": utils.Int64ToStr(ctx.Chat.Id),
	}

	dataGroup := map[string]string{
		"1": telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName),
		"2": utils.Int64ToStr(ctx.User.Id),
		"3": utils.IntToStr(0),
	}

	dataPrivate := map[string]string{
		"1": telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName),
		"2": utils.Int64ToStr(ctx.User.Id),
		"3": ctx.Chat.Title,
	}

	txtGroup, keybGroup := telegram.CreateMenuKeyboardf("./data/menu/username_group.json", 1, dataGroup, dataButton)
	txtPrivate, keybPrivate := telegram.CreateMenuKeyboardf("./data/menu/username_private.json", 1, dataPrivate, dataButton)
	txtLog := fmt.Sprintf(
		"#USERNAME"+
			"\n<b>User Name:</b> %s"+
			"\n<b>User ID:</b> <code>%v</code>"+
			"\n<b>Chat Name:</b> %s"+
			"\n<b>Chat ID:</b> <code>%v</code>"+
			"\n<b>Link:</b> %s",
		telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName),
		ctx.User.Id,
		ctx.Chat.Title,
		ctx.Chat.Id,
		telegram.CreateLinkHtml(telegram.CreateMessageLink(ctx.Chat, toDeleteAndSave), "Here"),
	)

	if !<-rstrChan {
		txtGroup += "\n\n🚫 <b>Tetapi saya tidak bisa membisukannya, mohon periksa kembali perizinan saya!</b>"
		ctx.SendMessage(txtGroup, 0)
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
		ctx.SendMessageKeyboard(txtGroup, 0, keybGroup)
		getPref.LastServiceMessageId = ctx.Message.MessageId
		_ = m.App.DB.Pref.SavePreference(getPref)
	}()

	go func() { defer wg.Done(); ctx.DeleteMessage(toDeleteAndSave) }()
	go func() { defer wg.Done(); ctx.SendMessageAsync(txtPrivate, ctx.User.Id, keybPrivate) }()
	go func() { defer wg.Done(); ctx.SendMessageAsync(txtLog, m.App.Config.LogEvent, nil) }()
	return telegram.EndOrder
}

func (m *Module) usernameCallbackGroup(ctx *telegram.TgContext) error {
	if telegram.IsPrivate(ctx.Chat.Type) {
		return ext.ContinueGroups
	}

	pattern, _ := regexp.Compile(`username\((.+?)\)\((.+?)\)`)
	userId := utils.StrToInt64(pattern.FindStringSubmatch(ctx.Callback.Data)[1])

	if !(userId == ctx.Callback.From.Id) {
		if ctx.User.Username == "" {
			ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
			return nil
		}

		ctx.UnRestrictMember(0, 0)
		ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
		return nil
	}

	if ctx.User.Username == "" {
		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
		return nil
	} else if ctx.User.Username != "" {
		ctx.UnRestrictMember(0, 0)
		ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
		ctx.DeleteMessage(0)
		return nil
	} else if p, err := ctx.Callback.From.GetProfilePhotos(ctx.Bot, nil); p != nil && p.TotalCount == 0 {
		if err != nil {
			ctx.AnswerCallback("Terjadi Kesalahan, Silahkan Coba Lagi", true)
			return err
		}

		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG FOTO PROFIL", true)
		return nil
	}

	ctx.UnRestrictMember(0, 0)
	ctx.AnswerCallback("✅ Terimakasih telah memasang Foto Profil", true)
	ctx.DeleteMessage(0)
	return telegram.ContinueOrder
}

func (m *Module) usernameCallbackPrivate(ctx *telegram.TgContext) error {
	if !telegram.IsPrivate(ctx.Chat.Type) {
		return ext.ContinueGroups
	}

	pattern, _ := regexp.Compile(`username\((.+?)\)\((.+?)\)`)
	userId := utils.StrToInt64(pattern.FindStringSubmatch(ctx.Callback.Data)[1])
	chatId := utils.StrToInt64(pattern.FindStringSubmatch(ctx.Callback.Data)[2])

	if !(userId == ctx.Callback.From.Id) {
		ctx.AnswerCallback("Anda dilarang menggunakan tombol ini!", true)
		return nil
	}

	if ctx.User.Username == "" {
		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
		return nil
	} else if ctx.User.Username != "" {
		ctx.UnRestrictMember(userId, chatId)
		ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
		ctx.DeleteMessage(0)
		return nil
	} else if p, err := ctx.Callback.From.GetProfilePhotos(ctx.Bot, nil); p != nil && p.TotalCount == 0 {
		if err != nil {
			ctx.AnswerCallback("Terjadi Kesalahan, Silahkan Coba Lagi", true)
			return err
		}

		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG FOTO PROFIL", true)
		return nil
	}

	ctx.UnRestrictMember(userId, chatId)
	ctx.AnswerCallback("✅ Terimakasih telah memasang Foto Profil/Username", true)
	ctx.DeleteMessage(0)
	return telegram.ContinueOrder
}
