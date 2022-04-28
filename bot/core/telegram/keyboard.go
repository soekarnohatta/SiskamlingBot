package telegram

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type Button struct {
	Text string
	Data string
}

// BuildKeyboard will handle buttons creation from a json file
// and defined size.
func BuildKeyboard(path string, size int) (res [][]gotgbot.InlineKeyboardButton) {
	openFile, err := os.Open(path)
	if err != nil {
		log.Print("failed to open file: " + err.Error())
		return
	}
	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {
			log.Print("failed to close file: " + err.Error())
			return
		}
	}(openFile)

	readFile, _ := ioutil.ReadAll(openFile)
	var result []Button
	_ = json.Unmarshal(readFile, &result)
	var btnList []gotgbot.InlineKeyboardButton

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

	return append(res, btnList)
}

// BuildKeyboardf will handle buttons creation from a json file
// and defined size but with extra args to be placed inside the button.
func BuildKeyboardf(path string, size int, dataMap map[string]string) (res [][]gotgbot.InlineKeyboardButton) {
	openFile, err := os.Open(path)
	if err != nil {
		log.Print("failed to open file: " + err.Error())
		return
	}
	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {
			log.Print("failed to close file: " + err.Error())
			return
		}
	}(openFile)

	readFile, _ := ioutil.ReadAll(openFile)
	var result []Button
	_ = json.Unmarshal(readFile, &result)

	var btnList []gotgbot.InlineKeyboardButton
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

	return append(res, btnList)
}

func isValidUrl(str string) bool {
	isValid, err := url.Parse(str)
	return err == nil && isValid.Scheme != "" && isValid.Host != ""
}
