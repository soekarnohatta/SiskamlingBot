package metrics

import (
	"SiskamlingBot/bot/models"
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
)

func UsernameMetrics(_ *gotgbot.Bot, ctx *ext.Context) error {
	err := models.SaveUser(context.TODO(), models.NewUser(
		ctx.Update.Message.From.Id,
		ctx.Update.Message.From.FirstName,
		ctx.Update.Message.From.LastName,
		ctx.Update.Message.From.Username,
	))
	if err != nil {
		log.Println("failed to update user due to: " + err.Error())
		return ext.ContinueGroups
	}

	return ext.ContinueGroups
}
