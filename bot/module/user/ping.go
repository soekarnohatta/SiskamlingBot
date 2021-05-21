package user

import (
	"SiskamlingBot/bot/core/telegram"
	"strconv"
)

func (m Module) ping(ctx *telegram.TgContext) {
	ctx.ReplyMessage("<b>Pong</b>")
	for x := range "loremipsumdolor" {
		ctx.EditMessage(strconv.Itoa(x))
	}

	ctx.EditMessage("Done")
}
