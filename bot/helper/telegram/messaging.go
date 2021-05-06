package telegram

import (
	"github.com/pkg/errors"
	"html"
	"io"
	"net/http"
	"os"
	"strconv"
)

// DownloadFile downloads file(s) from telegram servers.
func DownloadFile(telegramPath string, filePath string) (*os.File, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	get, errget := http.Get("https://api.telegram.org/file/bot" + "bot.Config.BotAPIKey" + "/" + telegramPath)
	if errget != nil {
		_ = os.Remove(filePath)
		return nil, err
	}

	if get != nil {
		defer get.Body.Close()
		if get.StatusCode != http.StatusOK {
			_ = os.Remove(filePath)
			return nil, errors.Wrapf(nil, "bad request: %v", get.StatusCode)
		}

		_, err := io.Copy(file, get.Body)
		if err != nil {
			_ = os.Remove(filePath)
			return nil, err
		}
	}

	return file, nil
}

// MentionHtml mentions a user by using HTML formatting.
func MentionHtml(userId int, name string) string {
	return "<a href=\"tg://user?id=" + strconv.Itoa(userId) + "\">" + html.EscapeString(name) + "</a>"
}

// CreateLinkHtml creates a link using HTML formatting.
func CreateLinkHtml(link string, txt string) string {
	return "<a href=\"" + link + "\">" + html.EscapeString(txt) + "</a>"
}
