package user

import (
	"SiskamlingBot/bot/core/telegram"
)

func (m Module) ping(ctx *telegram.TgContext) {
	ctx.ReplyMessage("<b>Ping</b>")
	ctx.EditMessage("<b>Pong</b>")
}
