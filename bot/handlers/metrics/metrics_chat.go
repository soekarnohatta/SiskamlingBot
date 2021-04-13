package metrics

import (
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/soekarnohatta/Siskamling/bot/models"
	"log"
)

func ChatMetrics(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.Update.Message == nil || ctx.Update.CallbackQuery == nil {
		return ext.EndGroups
	}

	err := models.SaveChat(context.TODO(), models.Chat{
		ChatID:    ctx.Update.Message.Chat.Id,
		ChatType:  ctx.Update.Message.Chat.Type,
		ChatLink:  ctx.Update.Message.Chat.InviteLink,
		ChatTitle: ctx.Update.Message.Chat.Title,
	})
	if err != nil {
		log.Println("failed to update chat due to: " + err.Error())
		return ext.ContinueGroups
	}

	return ext.ContinueGroups
}
