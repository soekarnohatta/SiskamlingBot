package core

import (
	"SiskamlingBot/bot/core/telegram"
)

func IsUserRestricted(ctx *telegram.TgContext) bool {
	getMember, _ := ctx.Bot.GetChatMember(ctx.Chat.Id, ctx.User.Id, nil)
	return getMember != nil && getMember.GetStatus() == "restricted"
}

func IsUserAdmin(ctx *telegram.TgContext) bool {
	member, err := ctx.Bot.GetChatMember(ctx.Chat.Id, ctx.User.Id, nil)
	if err != nil {
		return false
	}

	return member.MergeChatMember().Status == "administrator" || member.MergeChatMember().Status == "owner"
}
