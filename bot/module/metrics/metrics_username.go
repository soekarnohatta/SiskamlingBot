package metrics

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/model"
	"context"
	"log"
)

func (m Module) usernameMetric(ctx *telegram.TgContext) {
	err := model.SaveUser(m.Bot.DB, context.TODO(), model.NewUser(
		ctx.Message.From.Id,
		ctx.Message.From.FirstName,
		ctx.Message.From.LastName,
		ctx.Message.From.Username,
	))
	if err != nil {
		log.Println("failed to update user due to: " + err.Error())
		return
	}
}
