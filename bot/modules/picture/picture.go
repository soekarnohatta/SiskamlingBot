package picture

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"regexp"
	"sync"
)

func (m *Module) pictureScan(ctx *telegram.TgContext) error {
	var getPref, _ = m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if getPref != nil && !getPref.EnforcePicture {
		return telegram.ContinueOrder
	}

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)

	var rstrChan = make(chan bool, 1)
	var untilDate = utils.ExtractTime("5m")
	var toDeleteServiceMessage = getPref.LastServiceMessageId
	var toDeleteAndSave = ctx.Message.MessageId

	go func() { rstrChan <- ctx.RestrictMember(0, 0, untilDate) }()
	go func() { ctx.DeleteMessage(toDeleteServiceMessage); wg.Done() }()

	var dataButton = map[string]string{
		"1": utils.Int64ToStr(ctx.User.Id),
		"2": utils.Int64ToStr(ctx.Chat.Id),
	}

	var dataGroup = map[string]string{
		"1": telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName),
		"2": utils.Int64ToStr(ctx.User.Id),
		"3": utils.IntToStr(0),
	}

	var dataPrivate = map[string]string{
		"1": telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName),
		"2": utils.Int64ToStr(ctx.User.Id),
		"3": ctx.Chat.Title,
	}

	var txtGroup, keybGroup = telegram.CreateMenuKeyboardf("./data/menu/picture_group.json", 1, dataGroup, dataButton)
	var txtPrivate, keybPrivate = telegram.CreateMenuKeyboardf("./data/menu/picture_private.json", 1, dataPrivate, dataButton)
	var txtLog = fmt.Sprintf(
		"#PICTURE"+
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

		var err = m.App.DB.Pref.SavePreference(getPref)
		if err != nil {
			return err
		}
		return telegram.EndOrder
	}

	wg.Add(4)
	go func() {
		ctx.SendMessageKeyboard(txtGroup, 0, keybGroup)
		getPref.LastServiceMessageId = ctx.Message.MessageId
		var _ = m.App.DB.Pref.SavePreference(getPref)
		wg.Done()
	}()

	go func() { ctx.DeleteMessage(toDeleteAndSave); wg.Done() }()
	go func() { ctx.SendMessageAsync(txtPrivate, ctx.User.Id, keybPrivate); wg.Done() }()
	go func() { ctx.SendMessageAsync(txtLog, m.App.Config.LogEvent, nil); wg.Done() }()
	return telegram.EndOrder
}

func (m *Module) pictureCallbackGroup(ctx *telegram.TgContext) error {
	if telegram.IsPrivate(ctx.Chat.Type) {
		return ext.ContinueGroups
	}

	var pattern, _ = regexp.Compile(`picture\((.+?)\)\((.+?)\)`)
	var userId = utils.StrToInt64(pattern.FindStringSubmatch(ctx.Callback.Data)[1])

	if !(userId == ctx.Callback.From.Id) {
		if p, err := ctx.Callback.From.GetProfilePhotos(ctx.Bot, nil); p != nil && p.TotalCount == 0 {
			if err != nil {
				ctx.AnswerCallback("Terjadi Kesalahan, Silahkan Coba Lagi", true)
				return err
			}

			ctx.AnswerCallback("❌ ANDA BELUM MEMASANG FOTO PROFIL", true)
			return nil
		}

		ctx.UnRestrictMember(0, 0)
		ctx.AnswerCallback("✅ Terimakasih telah memasang Foto Profil", true)
		return nil
	}

	if p, err := ctx.Callback.From.GetProfilePhotos(ctx.Bot, nil); p != nil && p.TotalCount == 0 {
		if err != nil {
			ctx.AnswerCallback("Terjadi Kesalahan, Silahkan Coba Lagi", true)
			return err
		}

		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG FOTO PROFIL", true)
		return nil
	} else if ctx.User.Username == "" {
		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
		return nil
	} else if ctx.User.Username != "" {
		ctx.UnRestrictMember(0, 0)
		ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
		ctx.DeleteMessage(0)
		return nil
	}

	ctx.UnRestrictMember(0, 0)
	ctx.AnswerCallback("✅ Terimakasih telah memasang Foto Profil", true)
	ctx.DeleteMessage(0)
	return telegram.ContinueOrder
}

func (m Module) pictureCallbackPrivate(ctx *telegram.TgContext) error {
	if !telegram.IsPrivate(ctx.Chat.Type) {
		return ext.ContinueGroups
	}

	var pattern, _ = regexp.Compile(`picture\((.+?)\)\((.+?)\)`)
	var userId = utils.StrToInt64(pattern.FindStringSubmatch(ctx.Callback.Data)[1])
	var chatId = utils.StrToInt64(pattern.FindStringSubmatch(ctx.Callback.Data)[2])

	if !(userId == ctx.Callback.From.Id) {
		ctx.AnswerCallback("Anda dilarang menggunakan tombol ini!", true)
		return nil
	}

	if p, err := ctx.Callback.From.GetProfilePhotos(ctx.Bot, nil); p != nil && p.TotalCount == 0 {
		if err != nil {
			ctx.AnswerCallback("Terjadi Kesalahan, Silahkan Coba Lagi", true)
			return err
		}

		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG FOTO PROFIL", true)
		return nil
	} else if ctx.User.Username == "" {
		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
		return nil
	} else if ctx.User.Username != "" {
		ctx.UnRestrictMember(userId, chatId)
		ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
		ctx.DeleteMessage(0)
		return nil
	}

	ctx.UnRestrictMember(userId, chatId)
	ctx.AnswerCallback("✅ Terimakasih telah memasang Foto Profil", true)
	ctx.DeleteMessage(0)
	return telegram.ContinueOrder
}
