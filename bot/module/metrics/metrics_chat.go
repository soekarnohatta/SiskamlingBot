package metrics

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/model"
	"context"
	"log"
)

func (m Module) chatMetric(ctx *telegram.TgContext) {
	err := model.SaveChat(m.App.DB, context.TODO(), model.NewChat(
		ctx.Chat.Id,
		ctx.Chat.Type,
		ctx.Chat.InviteLink,
		ctx.Chat.Title,
	))
	if err != nil {
		log.Println("failed to update chat due to: " + err.Error())
		return
	}
}
