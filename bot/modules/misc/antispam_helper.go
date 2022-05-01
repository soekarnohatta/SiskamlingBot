package user

import (
	"SiskamlingBot/bot/utils"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/soekarnohatta/go-spamwatch/spamwatch"
)

var myClient = &http.Client{Timeout: 2 * time.Second}

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

func (m Module) isCASBan(userId int64) bool {
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

func (m Module) isSwBan(userId int64) bool {
	swClient, _ := spamwatch.NewClient("", m.App.Config.SWToken)
	ban, err := swClient.GetBan(int(userId))
	if err != nil {
		return false
	}

	return ban.Reason != ""
}

func (m Module) isLocalBan(userId int64) bool {
	getUser, _ := m.App.DB.User.GetUserById(userId)
	return getUser != nil && getUser.Gban
}

func (m Module) IsBan(userId int64) bool {
	// Add temporary fix regarding anonymous channel issue
	if userId == 136817688 || userId == 777000 {
		return false
	}

	CASChan := make(chan bool)
	SWChan := make(chan bool)
	LocalChan := make(chan bool)

	go func() { LocalChan <- m.isLocalBan(userId) }()
	go func() { CASChan <- m.isCASBan(userId) }()
	go func() { SWChan <- m.isSwBan(userId) }()

	select {
	default:
		return <-LocalChan || <-SWChan || <-CASChan
	case <-time.After(2 * time.Second):
		return false
	}
}
