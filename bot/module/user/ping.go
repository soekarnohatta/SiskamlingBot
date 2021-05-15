package user

import (
	"SiskamlingBot/bot/core/telegram"
	"strconv"
)

func (m Module) ping(ctx *telegram.TgContext) {
	ctx.ReplyMessage("<b>Pong</b>")
	for x, _ := range "loremipsumdolorsitamet" {
		ctx.EditMessage(strconv.Itoa(x))
	}

	ctx.EditMessage("Done")
	ctx.ReplyMessage("<b>Pong</b>")
}
