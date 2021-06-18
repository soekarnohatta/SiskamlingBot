package app

import (
	"SiskamlingBot/bot/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/soekarnohatta/go-spamwatch/spamwatch"
)

var (
	SWClient, _ = spamwatch.NewClient("", os.Getenv("SWTOKEN"))
	myClient    = &http.Client{Timeout: 5 * time.Second}
)

type (
	casBan struct {
		Ok     bool   `json:"ok"`
		Result result `json:"result"`
	}

	result struct {
		Offenses  int      `json:"offenses"`
		Messages  []string `json:"message"`
		TimeAdded int      `json:"time_added"`
	}
)

func IsCASBan(userId int64) bool {
	// Request data to CAS API.
	cas := "https://api.cas.chat/check?user_id=" + utils.IntToStr(int(userId))
	re, err := myClient.Get(cas)
	if err != nil {
		return false
	}
	defer re.Body.Close()

	// Deserialize it...
	var ban casBan
	err = json.NewDecoder(re.Body).Decode(&ban)
	//log.Printf("User %v is CAS Banned: %v", userId, ban.Ok)
	return ban.Ok
}

func IsSwBan(userId int64) bool {
	ban, err := SWClient.GetBan(int(userId))
	if err != nil {
		if err.Error() == "Token is invalid" {
			log.Fatalln(err.Error())
		}
		log.Println(err.Error())
	}

	//log.Printf("User %v is SW Banned: %v", userId, ban.Reason != "")
	return ban.Reason != ""
}

func IsBan(userId int64) bool {
	sendChan := make(chan bool, 2)

	go func() { sendChan <- IsCASBan(userId) }()
	go func() { sendChan <- IsSwBan(userId) }()

	select {
	case <-sendChan:
		return <-sendChan
	case <-time.After(5 * time.Second):
		return false
	}
}
