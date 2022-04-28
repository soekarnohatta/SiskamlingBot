package telegram

import (
	utils2 "SiskamlingBot/bot/utils"
	"html"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func MentionHtml(userId int, name string) string {
	return "<a href=\"tg://user?id=" + strconv.Itoa(userId) + "\">" + html.EscapeString(name) + "</a>"
}

func CreateLinkHtml(link, txt string) string {
	return "<a href=\"" + link + "\">" + html.EscapeString(txt) + "</a>"
}

func CreateMessageLink(chat *gotgbot.Chat, msgId int64) string {
	if chat.Username == "" {
		return "https://t.me/c/" + strings.TrimPrefix(utils2.IntToStr(int(chat.Id)), "-100") + "/" + utils2.IntToStr(int(msgId))
	}

	return "https://t.me/" + chat.Username + "/" + utils2.IntToStr(int(msgId))
}
