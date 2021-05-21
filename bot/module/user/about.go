package user

import (
	"SiskamlingBot/bot/core/telegram"	
	"SiskamlingBot/bot/util"

)

func (m Module) about(ctx *telegram.TgContext) {
	dataMap := map[string]string{"1": "My App", "2": m.App.Config.BotVer, "3": "@SoekarnoHatta"}
	text, keyb := util.CreateMenuf("./data/menu/about.json", 2, dataMap)
	ctx.ReplyMessageKeyboard(text, keyb)
}
