package metrics

import (
	"SiskamlingBot/bot/models"
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
)

func ChatMetrics(_ *gotgbot.Bot, ctx *ext.Context) error {
	err := models.SaveChat(context.TODO(), models.NewChat(
		ctx.Update.Message.Chat.Id,
		ctx.Update.Message.Chat.Type,
		ctx.Update.Message.Chat.InviteLink,
		ctx.Update.Message.Chat.Title,
	))
	if err != nil {
		log.Println("failed to update chat due to: " + err.Error())
		return ext.ContinueGroups
	}

	return ext.ContinueGroups
}
