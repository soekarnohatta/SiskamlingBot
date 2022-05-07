package metrics

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/models"
)

func (m Module) usernameMetric(ctx *telegram.TgContext) error {
	var getUser, err = m.App.DB.User.GetUserById(ctx.Message.From.Id)
	if getUser != nil {
		getUser.UserName = ctx.Message.From.Username
		getUser.FirstName = ctx.Message.From.FirstName
		getUser.LastName = ctx.Message.From.LastName
		var err = m.App.DB.User.SaveUser(getUser)
		if err != nil {
			return err
		}

		return telegram.ContinueOrder
	}

	var newUser = models.NewUser(
		ctx.Message.From.Id,
		ctx.Message.From.FirstName,
		ctx.Message.From.LastName,
		ctx.Message.From.Username,
		false,
	)

	err = m.App.DB.User.SaveUser(newUser)
	if err != nil {
		return err
	}

	return telegram.ContinueOrder
}

func (m Module) chatMetric(ctx *telegram.TgContext) error {
	var newChat = models.NewChat(
		ctx.Chat.Id,
		ctx.Chat.Type,
		ctx.Chat.InviteLink,
		ctx.Chat.Title,
	)

	var err = m.App.DB.Chat.SaveChat(newChat)
	if err != nil {
		return err
	}

	return telegram.ContinueOrder
}

func (m Module) preferenceMetric(ctx *telegram.TgContext) error {
	var getPref, err = m.App.DB.Pref.GetPreferenceById(ctx.Chat.Id)
	if getPref != nil {
		return telegram.ContinueOrder
	}

	var newPref = models.Preference{
		PreferenceID:         ctx.Chat.Id,
		EnforcePicture:       true,
		EnforceUsername:      true,
		EnforceAntispam:      true,
		LastServiceMessageId: 1,
	}

	err = m.App.DB.Pref.SavePreference(&newPref)
	if err != nil {
		return err
	}

	return telegram.ContinueOrder
}
