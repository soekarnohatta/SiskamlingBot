package telegram

import (
	"SiskamlingBot/bot/utils"
	"html"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func MentionHtml(userId int64, name string) string {
	return "<a href=\"tg://user?id=" + utils.Int64ToStr(userId) + "\">" + html.EscapeString(name) + "</a>"
}

func CreateLinkHtml(link, txt string) string {
	return "<a href=\"" + link + "\">" + html.EscapeString(txt) + "</a>"
}

func CreateMessageLink(chat *gotgbot.Chat, msgId int64) string {
	if chat.Username == "" {
		return "https://t.me/c/" + strings.TrimPrefix(utils.Int64ToStr(chat.Id), "-100") + "/" + utils.Int64ToStr(msgId)
	}

	return "https://t.me/" + chat.Username + "/" + utils.Int64ToStr(msgId)
}
