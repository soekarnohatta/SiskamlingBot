package helper

import (
	app "SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/model"
	"context"
)

func IsUserBotRestricted(ctx *telegram.TgContext, app *app.MyApp) bool {
	getUsername, _ := model.GetUsernameByID(app.DB, context.TODO(), ctx.Message.From.Id)
	if getUsername != nil && getUsername.ChatID == ctx.Message.Chat.Id && getUsername.IsMuted {
		if !IsUserRestricted(ctx) {return false}
		return true
	}

	getPicture, _ := model.GetPictureByID(app.DB, context.TODO(), ctx.User.Id)
	if getPicture != nil && getPicture.ChatID == ctx.Chat.Id && getPicture.IsMuted {
		if !IsUserRestricted(ctx) {return false}
		return true
	}

	return IsUserRestricted(ctx)
}

func IsUserRestricted(ctx *telegram.TgContext) bool {
	getMember, _ := ctx.Bot.GetChatMember(ctx.Chat.Id, ctx.User.Id)
	if getMember != nil && !getMember.CanSendMessages {return true}
	return false
}