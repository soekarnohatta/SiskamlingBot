package metrics

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/models"
)

func (m Module) usernameMetric(ctx *telegram.TgContext) {
	models.SaveUser(m.App.DB, models.NewUser(
		ctx.Message.From.Id,
		ctx.Message.From.FirstName,
		ctx.Message.From.LastName,
		ctx.Message.From.Username,
	))
	return
}

func (m Module) chatMetric(ctx *telegram.TgContext) {
	models.SaveChat(m.App.DB, models.NewChat(
		ctx.Chat.Id,
		ctx.Chat.Type,
		ctx.Chat.InviteLink,
		ctx.Chat.Title,
	))
	return
}