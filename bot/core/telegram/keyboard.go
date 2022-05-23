package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type Button struct {
	Text string
	Data string
}

func BuildKeyboard(path string, size int) ([][]gotgbot.InlineKeyboardButton, error) {
	openFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("BuildKeyboard: failed to open file to build new keyboard with error: %w", err)
	}

	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {
			return
		}
	}(openFile)

	readFile, _ := ioutil.ReadAll(openFile)
	var result []Button
	_ = json.Unmarshal(readFile, &result)
	var btnList []gotgbot.InlineKeyboardButton
	var res [][]gotgbot.InlineKeyboardButton

	for _, data := range result {
		if isValidUrl(data.Data) {
			btnList = append(btnList, gotgbot.InlineKeyboardButton{Text: data.Text, Url: data.Data})
		} else {
			btnList = append(btnList, gotgbot.InlineKeyboardButton{Text: data.Text, CallbackData: data.Data})
		}
	}

	for size < len(btnList) {
		btnList, res = btnList[size:], append(res, btnList[0:size:size])
	}

	return append(res, btnList), nil
}

func BuildKeyboardf(path string, size int, dataMap map[string]string) ([][]gotgbot.InlineKeyboardButton, error) {
	openFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("BuildKeyboardf: failed to open file to build new keyboard with error: %w", err)
	}
	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {
			return
		}
	}(openFile)

	readFile, _ := ioutil.ReadAll(openFile)
	var result []Button
	_ = json.Unmarshal(readFile, &result)

	var btnList []gotgbot.InlineKeyboardButton
	var res [][]gotgbot.InlineKeyboardButton
	for _, data := range result {
		var replData []string
		for k, v := range dataMap {
			replData = append(replData, "{"+k+"}", v)
		}
		newReplace := strings.NewReplacer(replData...)

		if isValidUrl(newReplace.Replace(data.Data)) {
			btnList = append(btnList, gotgbot.InlineKeyboardButton{
				Text: newReplace.Replace(data.Text),
				Url:  newReplace.Replace(data.Data),
			})
		} else {
			btnList = append(btnList, gotgbot.InlineKeyboardButton{
				Text:         newReplace.Replace(data.Text),
				CallbackData: newReplace.Replace(data.Data),
			})
		}
	}

	for size < len(btnList) {
		btnList, res = btnList[size:], append(res, btnList[0:size:size])
	}

	return append(res, btnList), nil
}

func isValidUrl(str string) bool {
	isValid, err := url.Parse(str)
	return err == nil && isValid.Scheme != "" && isValid.Host != ""
}
