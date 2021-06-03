package telegram

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
		log.Println("failed to open file: " + err.Error())
		return nil
	}
	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {
			log.Println("failed to close file: " + err.Error())
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
		return newMenu.Text, BuildKeyboard(newMenu.Keyboard, length)
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

		return newReplace.Replace(newMenu.Text), BuildKeyboardf(newMenu.Keyboard, length, dataMap)
	}
	return "", nil
}
