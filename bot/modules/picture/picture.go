package picture

import (
	"SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
	"fmt"
	"regexp"
	"strconv"
	"sync"
)

const (
	picLog = `#PICTURE
<b>User Name:</b> %s
<b>User ID:</b> <code>%v</code>
<b>Chat Name:</b> %s
<b>Chat ID:</b> <code>%v</code>
<b>Link:</b> %s`

	picMsg = "⚠ <b>%v</b> [<code>%v</code>] telah dibisukan karena belum memasang <b>Foto Profil!</b>"
)

func (m Module) pictureScan(ctx *telegram.TgContext) error {
	if core.IsUserRestricted(ctx) {
		return telegram.ContinueOrder
	}

	if !ctx.RestrictMember(0, 0) {
		unavailable := picMsg + "\n\n🚫 <b>Tetapi saya tidak bisa membisukannya, mohon periksa kembali perizinan saya!</b>"
		textToSend := fmt.Sprintf(unavailable, telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName), ctx.User.Id)
		ctx.SendMessage(textToSend, 0)
		return telegram.EndOrder
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() { defer wg.Done(); ctx.DeleteMessage(0) }()
	go func() {
		defer wg.Done()
		textToSend := fmt.Sprintf(picMsg, telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName), ctx.User.Id)
		keyboard, _ := telegram.BuildKeyboardf("./data/keyboard/picture.json", 1, map[string]string{"1": strconv.Itoa(int(ctx.User.Id))})
		ctx.SendMessageKeyboard(textToSend, 0, keyboard)
	}()
	go func() {
		defer wg.Done()
		textToSend := fmt.Sprintf(picLog,
			telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName),
			ctx.User.Id,
			ctx.Chat.Title,
			ctx.Chat.Id,
			telegram.CreateLinkHtml(telegram.CreateMessageLink(ctx.Chat, ctx.Message.MessageId), "Here"),
		)
		ctx.SendMessage(textToSend, m.App.Config.LogEvent)
	}()
	wg.Wait()
	return telegram.EndOrder
}

func (m Module) pictureCallback(ctx *telegram.TgContext) error {
	pattern, _ := regexp.Compile(`picture\((.+?)\)`)
	if !(pattern.FindStringSubmatch(ctx.Callback.Data)[1] == strconv.Itoa(int(ctx.Callback.From.Id))) {

		if p, err := ctx.Callback.From.GetProfilePhotos(ctx.Bot, nil); p != nil && p.TotalCount == 0 {
			if err != nil {
				ctx.AnswerCallback("Terjadi Kesalahan, Silahkan Coba Lagi", true)
				return err
			}

			ctx.AnswerCallback("❌ ANDA BELUM MEMASANG FOTO PROFIL", true)
			return nil
		}

		ctx.UnRestrictMember(0)
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
	}

	ctx.UnRestrictMember(0)
	ctx.AnswerCallback("✅ Terimakasih telah memasang Foto Profil", true)
	ctx.DeleteMessage(0)
	return telegram.ContinueOrder
}
