package user

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/util"

)

func (m Module) start(ctx *telegram.TgContext) {
	dataMap := map[string]string{"1": m.App.Bot.User.FirstName, "2": m.App.Config.BotVer, "3": "@SoekarnoHatta", "uname": m.App.Bot.User.Username}
	text, keyb := util.CreateMenuf("./data/menu/start.json", 2, dataMap)
	ctx.ReplyMessageKeyboard(text, keyb)
}
