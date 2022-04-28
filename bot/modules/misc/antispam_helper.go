package user

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/soekarnohatta/go-spamwatch/spamwatch"

	"SiskamlingBot/bot/utils"
)

var (
	swClient, _ = spamwatch.NewClient("", os.Getenv("SWTOKEN"))
	myClient    = &http.Client{Timeout: 2 * time.Second}
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

func isCASBan(userId int64) bool {
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

func isSwBan(userId int64) bool {
	ban, err := swClient.GetBan(int(userId))
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
	if userId == 136817688 || userId == 777000 {
		return false
	}

	CASChan := make(chan bool)
	SWChan := make(chan bool)

	go func() { CASChan <- isCASBan(userId) }()
	go func() { SWChan <- isSwBan(userId) }()

	select {
	default:
		return <-SWChan || <-CASChan
	case <-time.After(2 * time.Second):
		return false
	}
}
