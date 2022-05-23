package telegram

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type Menu struct {
	Callback string
	Keyboard string
	Text     string
}

func ParseMenu(path string) *Menu {
	openFile, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {
			return
		}
	}(openFile)

	readFile, _ := ioutil.ReadAll(openFile)
	var result *Menu
	_ = json.Unmarshal(readFile, &result)
	return result
}

func CreateMenu(path string, length int) (string, [][]gotgbot.InlineKeyboardButton) {
	newMenu := ParseMenu(path)
	if newMenu != nil {
		keyboard, _ := BuildKeyboard(newMenu.Keyboard, length)
		return newMenu.Text, keyboard
	}
	return "", nil
}

func CreateMenuf(path string, length int, dataMap map[string]string) (string, [][]gotgbot.InlineKeyboardButton) {
	newMenu := ParseMenu(path)
	if newMenu != nil {
		var replData []string
		for k, v := range dataMap {
			replData = append(replData, "{"+k+"}", v)
		}
		newReplace := strings.NewReplacer(replData...)
		keyboard, _ := BuildKeyboard(newMenu.Keyboard, length)
		return newReplace.Replace(newMenu.Text), keyboard
	}
	return "", nil
}

func CreateMenuKeyboardf(path string, length int, dataText, dataBtn map[string]string) (string, [][]gotgbot.InlineKeyboardButton) {
	newMenu := ParseMenu(path)
	if newMenu != nil {
		var replData []string
		for k, v := range dataText {
			replData = append(replData, "{"+k+"}", v)
		}
		newReplace := strings.NewReplacer(replData...)
		keyboard, _ := BuildKeyboardf(newMenu.Keyboard, length, dataBtn)
		return newReplace.Replace(newMenu.Text), keyboard
	}
	return "", nil
}
