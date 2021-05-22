package admin

import (
	"SiskamlingBot/bot/core/telegram"
	"encoding/json"
)

func (m Module) debug(ctx *telegram.TgContext) {
	if ctx.Message.ReplyToMessage != nil {
		output, _ := json.MarshalIndent(ctx.Message.ReplyToMessage, "", "  ")
		ctx.ReplyMessage(string(output))
		return
	}  

	output, _ := json.MarshalIndent(ctx.Message, "", "  ")
	ctx.ReplyMessage(string(output))
}