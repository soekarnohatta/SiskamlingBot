package helper

import (
	app "SiskamlingBot/bot/core"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/model"
	"context"
)

func IsUserBotRestricted(ctx *telegram.TgContext, app *app.MyApp) bool {
	if IsUsernameRestricted(ctx, app) {
		if !IsUserRestricted(ctx) {
			return false
		} else if IsProfileRestricted(ctx, app) {
			return true
		}
		return true
	} else if IsProfileRestricted(ctx, app) {
		if !IsUserRestricted(ctx) {
			return false
		} else if IsUsernameRestricted(ctx, app) {
			return true
		}
		return true
	}

	return IsUserRestricted(ctx)
}

func IsUsernameRestricted(ctx *telegram.TgContext, app *app.MyApp) bool {
	getUsername, _ := model.GetUsernameByID(app.DB, context.TODO(), ctx.Message.From.Id)
	return getUsername != nil && getUsername.ChatID == ctx.Message.Chat.Id && getUsername.IsMuted
}

func IsProfileRestricted(ctx *telegram.TgContext, app *app.MyApp) bool {
	getPicture, _ := model.GetPictureByID(app.DB, context.TODO(), ctx.User.Id)
	return getPicture != nil && getPicture.ChatID == ctx.Chat.Id && getPicture.IsMuted
}

func IsUserRestricted(ctx *telegram.TgContext) bool {
	getMember, _ := ctx.Bot.GetChatMember(ctx.Chat.Id, ctx.User.Id)
	if getMember != nil && !getMember.CanSendMessages {return true}
	return false
}