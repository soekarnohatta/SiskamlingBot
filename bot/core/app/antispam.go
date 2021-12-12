package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/soekarnohatta/go-spamwatch/spamwatch"

	"SiskamlingBot/bot/utils"
)

var (
	SWClient *spamwatch.Client
	myClient = &http.Client{Timeout: 5 * time.Second}
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
			log.Fatal(err.Error())
		}
		log.Print(err.Error())
	}

	//log.Printf("User %v is SW Banned: %v", userId, ban.Reason != "")
	return ban.Reason != ""
}

func IsBan(userId int64) bool {
	// Add temporary fix regarding anonymous channel issue
	if userId == 136817688 {
		return false
	}

	CASChan := make(chan bool)
	SWChan := make(chan bool)

	go func() { CASChan <- IsCASBan(userId) }()
	go func() { SWChan <- IsSwBan(userId) }()

	select {
	default:
		return <-SWChan || <-CASChan
	case <-time.After(3 * time.Second):
		return false
	}
}
