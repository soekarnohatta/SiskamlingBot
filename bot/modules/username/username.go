package username

import (
	"SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
	"fmt"
	"regexp"
	"strconv"
	"sync"
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

func (m Module) usernameScan(ctx *telegram.TgContext) error {
	getPref, err := m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if getPref != nil && !getPref.EnforcePicture {
		return telegram.ContinueOrder
	} else if err != nil {
		return err
	}

	if core.IsUserRestricted(ctx) {
		return telegram.ContinueOrder
	}

	var wgDel sync.WaitGroup
	wgDel.Add(1)
	go func() { defer wgDel.Done(); ctx.DeleteMessage(getPref.LastServiceMessageId) }()
	if !ctx.RestrictMember(0, 0) {
		wgDel.Wait()
		unavailable := unameMsg + "\n\n🚫 <b>Tetapi saya tidak bisa membisukannya, mohon periksa kembali perizinan saya!</b>"
		textToSend := fmt.Sprintf(unavailable, telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName), ctx.User.Id)
		ctx.SendMessage(textToSend, 0)
		getPref.LastServiceMessageId = ctx.Message.MessageId
		err := m.App.DB.Pref.SavePreference(getPref)
		if err != nil {
			return err
		}
		return telegram.EndOrder
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() { defer wg.Done(); ctx.DeleteMessage(0) }()
	go func() {
		defer wg.Done()
		textToSend := fmt.Sprintf(unameMsg, telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName), ctx.User.Id)
		keyboard, _ := telegram.BuildKeyboardf("./data/keyboard/username.json", 1, map[string]string{"1": strconv.Itoa(int(ctx.User.Id))})
		ctx.SendMessageKeyboard(textToSend, 0, keyboard)
		getPref.LastServiceMessageId = ctx.Message.MessageId
		_ = m.App.DB.Pref.SavePreference(getPref)
	}()
	go func() {
		defer wg.Done()
		textToSend := fmt.Sprintf(unameLog,
			telegram.MentionHtml(ctx.User.Id, ctx.User.FirstName),
			ctx.User.Id,
			ctx.Chat.Title,
			ctx.Chat.Id,
			telegram.CreateLinkHtml(telegram.CreateMessageLink(ctx.Chat, ctx.Message.MessageId), "Here"))
		ctx.SendMessage(textToSend, m.App.Config.LogEvent)
	}()
	wg.Wait()
	wgDel.Wait()
	return telegram.EndOrder
}

func (m Module) usernameCallback(ctx *telegram.TgContext) error {
	pattern, _ := regexp.Compile(`username\((.+?)\)`)
	if !(pattern.FindStringSubmatch(ctx.Callback.Data)[1] == strconv.Itoa(int(ctx.Callback.From.Id))) {

		if ctx.User.Username == "" {
			ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
			return nil
		}

		ctx.UnRestrictMember(0)
		ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
		return nil
	}

	if ctx.User.Username == "" {
		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
		return nil
	}

	ctx.UnRestrictMember(0)
	ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
	ctx.DeleteMessage(0)
	return telegram.ContinueOrder
}
