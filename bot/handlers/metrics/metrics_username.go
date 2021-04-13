package metrics

import (
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/soekarnohatta/Siskamling/bot/models"
	"log"
)

func UsernameMetrics(b *gotgbot.Bot, ctx *ext.Context) error {
	err := models.SaveUser(context.TODO(), models.User{
		UserID:    ctx.Update.Message.From.Id,
		FirstName: ctx.Update.Message.From.FirstName,
		LastName:  ctx.Update.Message.From.LastName,
		UserName:  ctx.Update.Message.From.Username,
	})
	if err != nil {
		log.Println("failed to update user due to: " + err.Error())
		return ext.ContinueGroups
	}

	return ext.ContinueGroups
}
