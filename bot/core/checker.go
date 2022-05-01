package core

import (
	"SiskamlingBot/bot/core/telegram"
)

func IsUserRestricted(ctx *telegram.TgContext) bool {
	getMember, _ := ctx.Bot.GetChatMember(ctx.Chat.Id, ctx.User.Id)
	if getMember != nil && getMember.GetStatus() == "restricted" {
		return true
	}
	return false
}
