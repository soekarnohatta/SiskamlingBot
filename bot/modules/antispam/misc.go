package user

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
	"fmt"
	"regexp"
	"time"

	"github.com/shirou/gopsutil/host"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

const infoTxt = "📁 <b>Bot Info</b>\n" +
	"<b>Bot Name :</b> <code>%v</code>\n" +
	"<b>Bot Username :</b> @%v\n" +
	"<b>Is Debug :</b> <code>%v</code>\n" +
	"<b>Version :</b> <code>%v</code>\n" +
	"<b>Bot Uptime :</b> <code>%v</code>\n\n" +
	"🖥️ <b>Platform Info</b>\n" +
	"<b>Host OS :</b> <code>%v</code>\n" +
	"<b>Host Name :</b> <code>%v</code>\n" +
	"<b>Host Uptime :</b> <code>%v</code>\n" +
	"<b>Kernel Version :</b> <code>%v</code>\n" +
	"<b>Platform :</b> <code>%v</code>\n" +
	"<b>Timestamp :</b> <code>%v</code>\n\n" +
	"👥 <b>%v</b>\n" +
	"<b>Chat Id :</b> <code>%v</code>\n" +
	"<b>Chat Type :</b> <code>%v</code>\n\n" +
	"👤 <b>%v</b>\n" +
	"<b>Is Sudo :</b> <code>%v</code>\n" +
	"<b>User Id :</b> <code>%v</code>"

func (m *Module) about(ctx *telegram.TgContext) error {
	var dataMap = map[string]string{
		"1": m.App.Bot.User.FirstName,
		"2": m.App.Config.BotVer,
		"3": "Devz",
	}

	var text, keyb = telegram.CreateMenuf("./data/menu/about.json", 2, dataMap)
	ctx.ReplyMessageKeyboard(text, keyb)
	return nil
}

func (*Module) ping(ctx *telegram.TgContext) error {
	var timeStart = time.Now()
	ctx.ReplyMessage("<b>Ping</b>")
	var timeEnd = time.Since(timeStart)
	var text = fmt.Sprintf("<b>Recorded Timespan Is:</b> <code>%v</code> ", timeEnd.String())
	ctx.EditMessage(text)
	return nil
}

func (m *Module) start(ctx *telegram.TgContext) error {
	var dataMap = map[string]string{
		"1":     m.App.Bot.User.FirstName,
		"2":     m.App.Config.BotVer,
		"3":     "Unknown",
		"uname": m.App.Bot.User.Username,
	}

	var text, keyb = telegram.CreateMenuf("./data/menu/start.json", 2, dataMap)
	ctx.ReplyMessageKeyboard(text, keyb)
	return nil
}

func (m *Module) info(ctx *telegram.TgContext) error {
	var info, _ = host.Info()
	var replyTxt = fmt.Sprintf(
		infoTxt,
		ctx.Bot.FirstName,
		ctx.Bot.Username,
		m.App.Config.IsDebug,
		m.App.Config.BotVer,
		utils.ConvertSeconds(uint64(time.Since(m.App.TimeStart).Seconds())),
		info.OS,
		info.Hostname,
		utils.ConvertSeconds(info.Uptime),
		info.KernelVersion,
		info.Platform,
		time.Now().Local(),
		ctx.Chat.Title,
		ctx.Chat.Id,
		ctx.Chat.Type,
		ctx.User.FirstName,
		telegram.IsSudo(ctx.User.Id, m.App.Config.SudoUsers),
		ctx.User.Id,
	)

	ctx.ReplyMessage(replyTxt)
	return nil
}

func (m *Module) helpCallback(ctx *telegram.TgContext) error {
	var pattern, _ = regexp.Compile(`help\((.+?)\)`)
	var match = pattern.FindStringSubmatch(ctx.Callback.Data)[1]

	switch match {
	case "main":
		var dataMap = map[string]string{
			"1": telegram.MentionHtml(ctx.Bot.User.Id, ctx.Bot.User.FirstName),
			"2": utils.IntToStr(int(ctx.Bot.User.Id)),
		}

		var text, keyb = telegram.CreateMenuf("./data/menu/help.json", 2, dataMap)
		var _, _, err = ctx.Callback.Message.EditText(
			ctx.Bot,
			text,
			&gotgbot.EditMessageTextOpts{
				ParseMode:   "HTML",
				ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb},
			})

		if err != nil {
			return err
		}
	case "about":
		var dataMap = map[string]string{
			"1": m.App.Bot.User.FirstName,
			"2": m.App.Config.BotVer,
			"3": "Devz",
		}

		var text, keyb = telegram.CreateMenuf("./data/menu/about.json", 2, dataMap)
		var _, _, err = ctx.Callback.Message.EditText(
			ctx.Bot,
			text,
			&gotgbot.EditMessageTextOpts{
				ParseMode:   "HTML",
				ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb},
			})

		if err != nil {
			return err
		}
	default:
		var text, keyb = telegram.CreateMenu("./data/menu/help/"+match+".json", 2)
		if text == "" {
			ctx.AnswerCallback("FITUR BELUM SIAP!", true)
		}

		var _, _, err = ctx.Callback.Message.EditText(
			ctx.Bot,
			text,
			&gotgbot.EditMessageTextOpts{
				ParseMode:   "HTML",
				ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb},
			},
		)

		if err != nil {
			return err
		}
	}

	return telegram.ContinueOrder
}
