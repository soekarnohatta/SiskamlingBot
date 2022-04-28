package picture

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/models"
	"fmt"
	"log"
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

func (m Module) pictureScan(ctx *telegram.TgContext) {
	// if core.IsUserRestricted(ctx) {
	// 	 return
	// }

	newPicture := models.NewPicture(ctx.User.Id, ctx.Chat.Id, true)
	models.SavePicture(m.App.DB, newPicture)

	if !ctx.RestrictMember(0, 0) {
		unavailable := picMsg + "\n\n🚫 <b>Tetapi saya tidak bisa membisukannya, mohon periksa kembali perizinan saya!</b>"
		textToSend := fmt.Sprintf(unavailable, telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName), ctx.User.Id)
		ctx.SendMessage(textToSend, 0)
		return
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() { defer wg.Done(); ctx.DeleteMessage(0) }()
	go func() {
		defer wg.Done()
		textToSend := fmt.Sprintf(picMsg, telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName), ctx.User.Id)
		ctx.SendMessageKeyboard(textToSend, 0, telegram.BuildKeyboardf("./data/keyboard/picture.json", 1, map[string]string{"1": strconv.Itoa(int(ctx.User.Id))}))
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
}

func (m Module) pictureCallback(ctx *telegram.TgContext) {
	pattern, _ := regexp.Compile(`picture\((.+?)\)`)
	if !(pattern.FindStringSubmatch(ctx.Callback.Data)[1] == strconv.Itoa(int(ctx.Callback.From.Id))) {
		getPicture := models.GetPictureByID(m.App.DB, ctx.Callback.From.Id)
		if getPicture != nil && getPicture.ChatID == ctx.Callback.Message.Chat.Id {
			if p, err := ctx.Callback.From.GetProfilePhotos(ctx.Bot, nil); p != nil && p.TotalCount == 0 {
				if err != nil {
					log.Print("failed to get pictures: " + err.Error())
					return
				}

				ctx.AnswerCallback("❌ ANDA BELUM MEMASANG FOTO PROFIL", true)
				return
			}

			models.DeletePictureByID(m.App.DB, ctx.Callback.From.Id)

			ctx.UnRestrictMember(0)
			ctx.AnswerCallback("✅ Terimakasih telah memasang Foto Profil", true)
			return
		}

		ctx.AnswerCallback("❌ ANDA BUKAN PENGGUNA YANG DIMAKSUD!", true)
		return
	}

	if p, err := ctx.Callback.From.GetProfilePhotos(ctx.Bot, nil); p != nil && p.TotalCount == 0 {
		if err != nil {
			log.Print("failed to get pictures: " + err.Error())
			return
		}

		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG FOTO PROFIL", true)
		return
	}

	models.DeletePictureByID(m.App.DB, ctx.Callback.From.Id)

	ctx.UnRestrictMember(0)
	ctx.AnswerCallback("✅ Terimakasih telah memasang Foto Profil", true)
	ctx.DeleteMessage(0)
}
