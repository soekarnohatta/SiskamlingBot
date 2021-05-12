package user

import (
	"SiskamlingBot/bot/core/telegram"
)

func (m *Module) ping(ctx *telegram.TgContext) {
	ctx.ReplyMessage("<b>Pong</b>")
}
