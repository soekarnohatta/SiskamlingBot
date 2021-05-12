package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"log"
	"strconv"
	"time"
)

/*
 * Message
 */

func (c *TgContext) SendMessage(text string, chatID int64) {
	c.TimeProc = time.Now().UTC().Sub(time.Unix(c.Message.Date, 0).UTC())
	text += "\n\n⏱ <code>" + strconv.FormatFloat(c.TimeInit.Seconds(), 'f', 3, 64) + " s</code> | ⌛ <code>" + strconv.FormatFloat(c.TimeProc.Seconds(), 'f', 3, 64) + " s</code>"

	if chatID != 0 {
		_, err := c.Bot.SendMessage(chatID, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	_, err := c.Bot.SendMessage(c.Chat.Id, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (c *TgContext) SendMessageKeyboard(text string, chatID int64, keyb [][]gotgbot.InlineKeyboardButton) {
	c.TimeProc = time.Now().UTC().Sub(time.Unix(c.Message.Date, 0).UTC())
	text += "\n\n⏱ <code>" + strconv.FormatFloat(c.TimeInit.Seconds(), 'f', 3, 64) + " s</code> | ⌛ <code>" + strconv.FormatFloat(c.TimeProc.Seconds(), 'f', 3, 64) + " s</code>"

	if chatID != 0 {
		_, err := c.Bot.SendMessage(chatID, text, &gotgbot.SendMessageOpts{ParseMode: "HTML", ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb}})
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	_, err := c.Bot.SendMessage(c.Chat.Id, text, &gotgbot.SendMessageOpts{ParseMode: "HTML", ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb}})
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func (c *TgContext) ReplyMessage(text string) {
	c.TimeProc = time.Now().UTC().Sub(time.Unix(c.Message.Date, 0).UTC())
	text += "\n\n⏱ <code>" + strconv.FormatFloat(c.TimeInit.Seconds(), 'f', 3, 64) + " s</code> | ⌛ <code>" + strconv.FormatFloat(c.TimeProc.Seconds(), 'f', 3, 64) + " s</code>"

	_, err := c.Context.EffectiveMessage.Reply(c.Bot, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (c *TgContext) ReplyMessageKeyboard(text string, keyb [][]gotgbot.InlineKeyboardButton) {
	c.TimeProc = time.Now().UTC().Sub(time.Unix(c.Message.Date, 0).UTC())
	text += "\n\n⏱ <code>" + strconv.FormatFloat(c.TimeInit.Seconds(), 'f', 3, 64) + " s</code> | ⌛ <code>" + strconv.FormatFloat(c.TimeProc.Seconds(), 'f', 3, 64) + " s</code>"

	_, err := c.Message.Reply(c.Bot, text, &gotgbot.SendMessageOpts{ParseMode: "HTML", ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb}})
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (c *TgContext) DeleteMessage(msgId int64) {
	if msgId != 0 {
		_, err := c.Bot.DeleteMessage(c.Chat.Id, msgId)
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	_, err := c.Bot.DeleteMessage(c.Chat.Id, c.Message.MessageId)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

/*
 * Callback
 */

func (c *TgContext) AnswerCallback(text string, alert bool) {
	_, err := c.Callback.Answer(c.Bot, &gotgbot.AnswerCallbackQueryOpts{
		Text:      text,
		ShowAlert: alert,
		//CacheTime: 0,
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}
