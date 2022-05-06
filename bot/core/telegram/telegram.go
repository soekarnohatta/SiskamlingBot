package telegram

import (
	"strconv"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

var (
	defaultParseMode = &gotgbot.SendMessageOpts{
		ParseMode:                "HTML",
		AllowSendingWithoutReply: true,
		DisableWebPagePreview:    true,
		ReplyToMessageId:         -1,
	}

	replyDefaultParseMode = &gotgbot.SendMessageOpts{
		ParseMode:                "HTML",
		AllowSendingWithoutReply: true,
		DisableWebPagePreview:    true,
	}
)

/*
 * Message
 */

func (c *TgContext) SendMessage(text string, chatID int64) {
	if text == "" {
		text = "Bad Request: No text supplied!"
	}

	timeProc := strconv.FormatFloat(time.Since(time.Unix(c.Date, 0)).Seconds(), 'f', 3, 64)
	text += "\n\n⏱ <code>" + c.TimeInit + " s</code> | ⌛ <code>" + timeProc + " s</code>"

	if chatID != 0 {
		msg, err := c.Bot.SendMessage(chatID, text, defaultParseMode)
		if err != nil {
			return
		}

		c.Message = msg
		return
	}

	msg, err := c.Bot.SendMessage(c.Chat.Id, text, defaultParseMode)
	if err != nil {
		return
	}

	c.Message = msg
}

func (c *TgContext) SendMessageKeyboard(text string, chatID int64, keyb [][]gotgbot.InlineKeyboardButton) {
	if text == "" {
		text = "Bad Request: No text supplied!"
	}

	timeProc := strconv.FormatFloat(time.Since(time.Unix(c.Date, 0)).Seconds(), 'f', 3, 64)
	text += "\n\n⏱ <code>" + c.TimeInit + " s</code> | ⌛ <code>" + timeProc + " s</code>"
	msgOpt := &gotgbot.SendMessageOpts{
		ParseMode:   "HTML",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb},
	}

	if chatID != 0 {
		msg, err := c.Bot.SendMessage(chatID, text, msgOpt)
		if err != nil {
			return
		}

		c.Message = msg
		return
	}

	msg, err := c.Bot.SendMessage(c.Chat.Id, text, msgOpt)
	if err != nil {
		return
	}
	c.Message = msg
}

func (c *TgContext) ReplyMessage(text string) {
	if text == "" {
		text = "Bad Request: No text supplied!"
	}

	timeProc := strconv.FormatFloat(time.Since(time.Unix(c.Date, 0)).Seconds(), 'f', 3, 64)
	text += "\n\n⏱ <code>" + c.TimeInit + " s</code> | ⌛ <code>" + timeProc + " s</code>"
	replyDefaultParseMode.ReplyToMessageId = c.Message.MessageId

	msg, err := c.Bot.SendMessage(c.Chat.Id, text, replyDefaultParseMode)
	if err != nil {
		c.SendMessage(text, 0)
		return
	}

	c.Message = msg
}

func (c *TgContext) ReplyMessageKeyboard(text string, keyb [][]gotgbot.InlineKeyboardButton) {
	if text == "" {
		text = "Bad Request: No text supplied!"
	}

	timeProc := strconv.FormatFloat(time.Since(time.Unix(c.Date, 0)).Seconds(), 'f', 3, 64)
	text += "\n\n⏱ <code>" + c.TimeInit + " s</code> | ⌛ <code>" + timeProc + " s</code>"
	msgOpt := &gotgbot.SendMessageOpts{
		ParseMode:        "HTML",
		ReplyMarkup:      gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb},
		ReplyToMessageId: c.Message.MessageId,
	}

	msg, err := c.Bot.SendMessage(c.Chat.Id, text, msgOpt)
	if err != nil {
		return
	}

	c.Message = msg
}

func (c *TgContext) EditMessage(text string) {
	if text == "" {
		text = "Bad Request: No text supplied!"
	}

	timeProc := strconv.FormatFloat(time.Since(time.Unix(c.Date, 0)).Seconds(), 'f', 3, 64)
	text += "\n\n⏱ <code>" + c.TimeInit + " s</code> | ⌛ <code>" + timeProc + " s</code>"

	msg, _, err := c.Message.EditText(c.Bot, text, &gotgbot.EditMessageTextOpts{ParseMode: "HTML"})
	if err != nil {
		return
	}

	c.Message = msg
}

func (c *TgContext) DeleteMessage(msgId int64) {
	if msgId != 0 {
		_, err := c.Bot.DeleteMessage(c.Chat.Id, msgId, nil)
		if err != nil {
			return
		}
		return
	}

	_, err := c.Bot.DeleteMessage(c.Chat.Id, c.Message.MessageId, nil)
	if err != nil {
		return
	}
}

/*
 * Callback
 */

func (c *TgContext) AnswerCallback(text string, alert bool) {
	newAnswerCallbackQueryOpts := &gotgbot.AnswerCallbackQueryOpts{
		Text:      text,
		ShowAlert: alert,
	}

	_, err := c.Callback.Answer(c.Bot, newAnswerCallbackQueryOpts)
	if err != nil {
		return
	}
}

/*
 * ChatMember
 */

func (c *TgContext) RestrictMember(userId, untilDate int64) bool {
	if userId == 0 {
		userId = c.User.Id
	}

	if untilDate == 0 {
		untilDate = -1
	}

	newOpt := &gotgbot.RestrictChatMemberOpts{UntilDate: untilDate}
	newChatPermission := gotgbot.ChatPermissions{
		CanSendMessages:      false,
		CanSendMediaMessages: false,
		CanSendPolls:         false,
		CanSendOtherMessages: false,
	}

	_, err := c.Bot.RestrictChatMember(c.Chat.Id, userId, newChatPermission, newOpt)
	return err == nil
}

func (c *TgContext) UnRestrictMember(userId int64) bool {
	if userId == 0 {
		userId = c.User.Id
	}

	newOpt := &gotgbot.RestrictChatMemberOpts{UntilDate: -1}
	newChatPermission := gotgbot.ChatPermissions{
		CanSendMessages:      true,
		CanSendMediaMessages: true,
		CanSendPolls:         true,
		CanSendOtherMessages: true,
	}

	_, err := c.Bot.RestrictChatMember(c.Chat.Id, userId, newChatPermission, newOpt)
	return err == nil
}

func (c *TgContext) BanChatMember(userId int64) bool {
	if userId == 0 {
		userId = c.User.Id
	}

	banOpt := &gotgbot.BanChatMemberOpts{
		UntilDate:      0,
		RevokeMessages: true,
		RequestOpts:    nil,
	}

	_, err := c.Bot.BanChatMember(c.Chat.Id, userId, banOpt)
	if err != nil {
		return false
	}

	return true
}

func (c *TgContext) UnBanChatMember(userId int64) bool {
	if userId == 0 {
		userId = c.User.Id
	}

	unbanOpt := &gotgbot.UnbanChatMemberOpts{
		OnlyIfBanned: true,
		RequestOpts:  nil,
	}

	_, err := c.Bot.UnbanChatMember(c.Chat.Id, userId, unbanOpt)
	if err != nil {
		return false
	}

	return true
}
