package app

import (
	"SiskamlingBot/bot/util"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/soekarnohatta/go-spamwatch/spamwatch"
)

var (
	SWClient, _ = spamwatch.NewClient("", getSWToken())
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

func getSWToken() string {
	return os.Getenv("SWTOKEN")
}

func IsCASBan(userId int64) bool {
	// Request data to CAS API.
	cas := "https://api.cas.chat/check?user_id=" + util.IntToStr(int(userId))
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
		log.Fatalln(err.Error())
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
