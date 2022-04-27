package app

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/soekarnohatta/go-spamwatch/spamwatch"

	"SiskamlingBot/bot/utils"
)

var (
	SWClient *spamwatch.Client
	myClient = &http.Client{Timeout: 2 * time.Second}
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(re.Body)

	// Deserialize it...
	var ban casBan
	err = json.NewDecoder(re.Body).Decode(&ban)
	return ban.Ok
}

func IsSwBan(userId int64) bool {
	ban, err := SWClient.GetBan(int(userId))
	if err != nil {
		if err.Error() == "Token is invalid" {
			log.Fatal(err.Error())
			return false
		}
		log.Print(err.Error())
		return false
	}

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
	case <-time.After(2 * time.Second):
		return false
	}
}
