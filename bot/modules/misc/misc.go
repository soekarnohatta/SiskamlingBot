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

func (m Module) about(ctx *telegram.TgContext) error {
	dataMap := map[string]string{"1": m.App.Bot.User.FirstName, "2": m.App.Config.BotVer, "3": "Unknown"}
	text, keyb := telegram.CreateMenuf("./data/menu/about.json", 2, dataMap)
	ctx.ReplyMessageKeyboard(text, keyb)
	return nil
}

func (Module) ping(ctx *telegram.TgContext) error {
	timeStart := time.Now()
	ctx.ReplyMessage("<b>Ping</b>")
	timeEnd := time.Since(timeStart)
	ctx.EditMessage(timeEnd.String())
	return nil
}

func (m Module) start(ctx *telegram.TgContext) error {
	dataMap := map[string]string{"1": m.App.Bot.User.FirstName, "2": m.App.Config.BotVer, "3": "Unknown", "uname": m.App.Bot.User.Username}
	text, keyb := telegram.CreateMenuf("./data/menu/start.json", 2, dataMap)
	ctx.ReplyMessageKeyboard(text, keyb)
	return nil
}

func (m Module) info(ctx *telegram.TgContext) error {
	info, _ := host.Info()
	replyTxt := fmt.Sprintf("<b>Bot Info</b>\n"+
		"  👤 <b>Bot Name :</b> <code>%v</code>\n"+
		"  🤖 <b>Bot Username :</b> @%v\n"+
		"  🖥 <b>Host OS :</b> <code>%v</code>\n"+
		"  ⚙ <b>Host Name :</b> <code>%v</code>\n"+
		"  ⏱ <b>Host Uptime :</b> <code>%v</code>\n"+
		"  💽 <b>Kernel Version :</b> <code>%v</code>\n"+
		"  💾 <b>Platform :</b> <code>%v</code>\n"+
		"  📅 <b>Timestamp :</b> <code>%v</code>\n",
		ctx.Bot.FirstName,
		ctx.Bot.Username,
		info.OS,
		info.Hostname,
		utils.ConvertSeconds(info.Uptime),
		info.KernelVersion,
		info.Platform,
		time.Now().Local(),
	)
	ctx.ReplyMessage(replyTxt)
	return nil
}

func (Module) helpCallback(ctx *telegram.TgContext) error {
	pattern, _ := regexp.Compile(`help\((.+?)\)`)
	switch pattern.FindStringSubmatch(ctx.Callback.Data)[1] {
	case "main":
		dataMap := map[string]string{"1": telegram.MentionHtml(int(ctx.Bot.User.Id), ctx.Bot.User.FirstName), "2": utils.IntToStr(int(ctx.Bot.User.Id))}
		text, keyb := telegram.CreateMenuf("./data/menu/help.json", 2, dataMap)
		_, _, err := ctx.Callback.Message.EditText(ctx.Bot, text, &gotgbot.EditMessageTextOpts{ParseMode: "HTML", ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb}})
		if err != nil {
			return err
		}
	default:
		ctx.AnswerCallback("FITUR BELUM SIAP!", true)
	}
	return telegram.ContinueOrder
}
