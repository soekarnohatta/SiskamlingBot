package username

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/soekarnohatta/Siskamling/bot"
	"github.com/soekarnohatta/Siskamling/bot/helpers/telegram"
	"log"
	"regexp"
	"strconv"
)

func Username(b *gotgbot.Bot, ctx *ext.Context) error {
	if !telegram.IsGroup(ctx.Message.Chat.Type) {
		return nil
	}

	if ctx.Message.From.Username != "" || ctx.Message.From.Id == 777000{
		return nil
	}

	_, err := b.RestrictChatMember(ctx.Message.Chat.Id, ctx.Message.From.Id, gotgbot.ChatPermissions{
		CanSendMessages:       false,
		CanSendMediaMessages:  false,
		CanSendPolls:          false,
		CanSendOtherMessages:  false,
	},
	&gotgbot.RestrictChatMemberOpts{UntilDate: -1},
	)
	if err != nil {
		log.Println("failed to restrict member: " + err.Error())
		return nil
	}

	_, err = b.DeleteMessage(ctx.Message.Chat.Id, ctx.Message.MessageId)
	if err != nil {
		log.Println("failed to delete message: " + err.Error())
		return nil
	}

	textToSend := fmt.Sprintf("⚠ Pengguna <b>%v</b> [<code>%v</code>] telah dibisukan karena belum memasang <b>Username!</b>",
		telegram.MentionHtml(int(ctx.Message.From.Id), ctx.Message.From.FirstName),
		ctx.Message.From.Id,
		)

	_, err = b.SendMessage(ctx.Message.Chat.Id, textToSend, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: telegram.BuildKeyboardf("data/keyboard/username.json", 1,
				map[string]string{"1": strconv.Itoa(int(ctx.Message.From.Id))},
			),
		}})
	if err != nil {
		log.Println("failed to send message: " + err.Error())
		return nil
	}

	err = logusername(b, ctx)
	if err != nil {
		log.Println("failed to send log message: " + err.Error())
		return ext.ContinueGroups
	}

	return nil
}

func UsernameCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	pattern, _ := regexp.Compile(`username\((.+?)\)`)

	if !pattern.MatchString(cb.Data) {
		return nil
	}

	if !(pattern.FindStringSubmatch(cb.Data)[1] == strconv.Itoa(int(cb.From.Id))) {
		_, err := cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "❌ ANDA BUKAN PENGGUNA YANG DIMAKSUD!",
			ShowAlert: true,
			CacheTime: 0,
		})
		if err != nil {
			log.Println("failed to answer callbackquery: " + err.Error())
			return nil
		}
		return nil
	}

	if cb.From.Username == "" {
		_, err := cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "❌ Anda belum memasang username",
			ShowAlert: true,
			CacheTime: 0,
		})
		if err != nil {
			log.Println("failed to answer callbackquery: " + err.Error())
			return nil
		}
		return nil
	}

	_, err := cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text:      "✅ Terimakasih telah memasang username",
		ShowAlert: true,
		CacheTime: 0,
	})
	if err != nil {
		log.Println("failed to answer callbackquery: " + err.Error())
		return nil
	}

	_, err = cb.Message.Delete(b)
	if err != nil {
		log.Println("failed to delete message: " + err.Error())
		return nil
	}

	_, err = b.RestrictChatMember(cb.Message.Chat.Id, cb.From.Id, gotgbot.ChatPermissions{
		CanSendMessages:       true,
		CanSendMediaMessages:  true,
		CanSendPolls:          true,
		CanSendOtherMessages:  true,
	}, nil)
	if err != nil {
		log.Println("failed to restrict chatmember: " + err.Error())
		return nil
	}

	return nil
}

func logusername(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.Update.Message.From
	chat := ctx.Update.Message.Chat
	textToSend := fmt.Sprintf(`#USERNAME
<b>User Name:</b> %v
<b>User ID:</b> <code>%v</code>
<b>Chat Name:</b> %v
<b>Chat ID:</b> <code>%v</code>
<b>Link:</b> %v`,
telegram.MentionHtml(int(user.Id), user.FirstName), user.Id,
chat.Title, chat.Id,
telegram.CreateLinkHtml("https://t.me/" + chat.Username + "/" + strconv.Itoa(int(ctx.Update.Message.MessageId)), "Here"))

	_, err := b.SendMessage(bot.Config.LogEvent, textToSend, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	return err
}