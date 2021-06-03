package user

import (
	"SiskamlingBot/bot/core/telegram"
)

func (m Module) about(ctx *telegram.TgContext) {
	dataMap := map[string]string{"1": m.App.Bot.User.FirstName, "2": m.App.Config.BotVer, "3": "Unknown"}
	text, keyb := telegram.CreateMenuf("./data/menu/about.json", 2, dataMap)
	ctx.ReplyMessageKeyboard(text, keyb)
}
