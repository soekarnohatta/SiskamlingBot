package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
	_ = json.Unmarshal(readFile, result)
	return result
}

func CreateMenu() {
	files, err := ioutil.ReadDir("./data/menu")
	if err != nil {
		log.Println("failed to read dir: " + err.Error())
		return
	}

	var menuList []Menu
	for _, data := range files {
		newMenu := ParseMenu(data.Name())
		menuList = append(menuList, *newMenu)
	}
}
