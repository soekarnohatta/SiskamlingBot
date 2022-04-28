package core

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/models"
)

func IsUserBotRestricted(ctx *telegram.TgContext, app *app.MyApp) bool {
	if IsUsernameRestricted(ctx, app) {
		if IsUserRestricted(ctx) {
			return true
		} else if IsProfileRestricted(ctx, app) {
			return true
		} else {
			return true
		}
	} else if IsProfileRestricted(ctx, app) {
		if IsUserRestricted(ctx) {
			return true
		} else if IsUsernameRestricted(ctx, app) {
			return true
		} else {
			return true
		}
	}
	return IsUserRestricted(ctx)
}

func IsUsernameRestricted(ctx *telegram.TgContext, app *app.MyApp) bool {
	getUsername := models.GetUsernameByID(app.DB, ctx.Message.From.Id)
	return getUsername != nil && getUsername.ChatID == ctx.Message.Chat.Id && getUsername.IsMuted
}

func IsProfileRestricted(ctx *telegram.TgContext, app *app.MyApp) bool {
	getPicture := models.GetPictureByID(app.DB, ctx.User.Id)
	return getPicture != nil && getPicture.ChatID == ctx.Chat.Id && getPicture.IsMuted
}

func IsUserRestricted(ctx *telegram.TgContext) bool {
	getMember, _ := ctx.Bot.GetChatMember(ctx.Chat.Id, ctx.User.Id)
	if getMember != nil && getMember.GetStatus() == "restricted" {
		return true
	}
	return false
}
