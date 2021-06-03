package user

import (
	"SiskamlingBot/bot/core/telegram"
)

func (m Module) start(ctx *telegram.TgContext) {
	dataMap := map[string]string{"1": m.App.Bot.User.FirstName, "2": m.App.Config.BotVer, "3": "Unknown", "uname": m.App.Bot.User.Username}
	text, keyb := telegram.CreateMenuf("./data/menu/start.json", 2, dataMap)
	ctx.ReplyMessageKeyboard(text, keyb)
}
