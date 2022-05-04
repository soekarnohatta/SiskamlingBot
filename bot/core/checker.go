package core

import (
	"SiskamlingBot/bot/core/telegram"
)

func IsUserRestricted(ctx *telegram.TgContext) bool {
	getMember, _ := ctx.Bot.GetChatMember(ctx.Chat.Id, ctx.User.Id, nil)
	return getMember != nil && getMember.GetStatus() == "restricted"
}
