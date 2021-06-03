package user

import (
	"SiskamlingBot/bot/core/telegram"

	"time"
)

func (m Module) ping(ctx *telegram.TgContext) {
	timeStart := time.Now()
	ctx.ReplyMessage("<b>Ping</b>")
	timeEnd := time.Since(timeStart)
	ctx.EditMessage(timeEnd.String())
}
