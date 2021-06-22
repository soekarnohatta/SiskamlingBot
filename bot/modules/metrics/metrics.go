package metrics

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/models"
	"log"
)

func (m Module) usernameMetric(ctx *telegram.TgContext) {
	err := models.SaveUser(m.App.DB, models.NewUser(
		ctx.Message.From.Id,
		ctx.Message.From.FirstName,
		ctx.Message.From.LastName,
		ctx.Message.From.Username,
	))

	if err != nil {
		log.Println("failed to update user due to: " + err.Error())
	}
}

func (m Module) chatMetric(ctx *telegram.TgContext) {
	err := models.SaveChat(m.App.DB, models.NewChat(
		ctx.Chat.Id,
		ctx.Chat.Type,
		ctx.Chat.InviteLink,
		ctx.Chat.Title,
	))

	if err != nil {
		log.Println("failed to update chat due to: " + err.Error())
	}
}
