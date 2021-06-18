package user

import (
	"SiskamlingBot/bot/core/telegram"
	"time"
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

