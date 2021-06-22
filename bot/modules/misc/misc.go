package user

import (
	"SiskamlingBot/bot/core/telegram"
	"regexp"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func (m Module) about(ctx *telegram.TgContext) {
	dataMap := map[string]string{"1": m.App.Bot.User.FirstName, "2": m.App.Config.BotVer, "3": "Unknown"}
	text, keyb := telegram.CreateMenuf("./data/menu/about.json", 2, dataMap)
	ctx.ReplyMessageKeyboard(text, keyb)
}

func (m Module) ping(ctx *telegram.TgContext) {
	timeStart := time.Now()
	ctx.ReplyMessage("<b>Ping</b>")
	timeEnd := time.Since(timeStart)
	ctx.EditMessage(timeEnd.String())
}

func (m Module) start(ctx *telegram.TgContext) {
	dataMap := map[string]string{"1": m.App.Bot.User.FirstName, "2": m.App.Config.BotVer, "3": "Unknown", "uname": m.App.Bot.User.Username}
	text, keyb := telegram.CreateMenuf("./data/menu/start.json", 2, dataMap)
	ctx.ReplyMessageKeyboard(text, keyb)
}

func (m Module) helpCallback(ctx *telegram.TgContext) {
	pattern, _ := regexp.Compile(`help\((.+?)\)`)
	switch pattern.FindStringSubmatch(ctx.Callback.Data)[1] {
	case "main":
		text, keyb := telegram.CreateMenu("./data/menu/help.json", 2)
		ctx.Callback.Message.EditText(ctx.Bot, text, &gotgbot.EditMessageTextOpts{ParseMode: "HTML", ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb}})
	default :
		ctx.AnswerCallback("FITUR BELUM SIAP!", true)
	}
}
