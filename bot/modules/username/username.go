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
	if core.IsUserRestricted(ctx) {
		return telegram.ContinueOrder
	}

	if !ctx.RestrictMember(0, 0) {
		unavailable := unameMsg + "\n\n🚫 <b>Tetapi saya tidak bisa membisukannya, mohon periksa kembali perizinan saya!</b>"
		textToSend := fmt.Sprintf(unavailable, telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName), ctx.User.Id)
		ctx.SendMessage(textToSend, 0)
		return telegram.EndOrder
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() { defer wg.Done(); ctx.DeleteMessage(0) }()

	go func() {
		defer wg.Done()
		textToSend := fmt.Sprintf(unameMsg, telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName), ctx.User.Id)
		keyboard, _ := telegram.BuildKeyboardf("./data/keyboard/username.json", 1, map[string]string{"1": strconv.Itoa(int(ctx.User.Id))})
		ctx.SendMessageKeyboard(textToSend, 0, keyboard)
	}()

	go func() {
		defer wg.Done()
		textToSend := fmt.Sprintf(unameLog,
			telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName),
			ctx.User.Id,
			ctx.Chat.Title,
			ctx.Chat.Id,
			telegram.CreateLinkHtml(telegram.CreateMessageLink(ctx.Chat, ctx.Message.MessageId), "Here"))

		ctx.SendMessage(textToSend, m.App.Config.LogEvent)
	}()
	wg.Wait()
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
