package user

import (
	"SiskamlingBot/bot/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type (
	casBan struct {
		Ok     bool   `json:"ok"`
		Result result `json:"result"`
	}

	result struct {
		Offenses  int      `json:"offenses"`
		TimeAdded int      `json:"time_added"`
		Messages  []string `json:"message"`
	}
)

type banList struct {
	Admin   int    `json:"admin,omitempty"`
	Date    int64  `json:"date,omitempty"`
	Id      int64  `json:"id"`
	Reason  string `json:"reason"`
	Message string `json:"message,omitempty"`
}

func (m *Module) isCASBan(userId int64) bool {
	// Request data to CAS API.
	cas := "https://api.cas.chat/check?user_id=" + utils.IntToStr(int(userId))
	re, err := http.Get(cas)
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

func (m *Module) isSwBan(userId int64) bool {
	r, err := http.NewRequest(http.MethodGet, "https://api.spamwat.ch/banlist/"+utils.Int64ToStr(userId), nil)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.App.Config.SWToken))
	r.Header.Set("Content-Type", "application/json")

	if err != nil {
		return false
	}

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Do(r)
	if err != nil {
		return false
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	swBan := &banList{}
	err = json.NewDecoder(resp.Body).Decode(&swBan)
	if err != nil {
		return false
	}

	return swBan.Reason != "" || resp.StatusCode == http.StatusOK
}

func (m *Module) isLocalBan(userId int64) bool {
	getUser, _ := m.App.DB.User.GetUserById(userId)
	return getUser != nil && getUser.Gban
}

func (m *Module) IsBan(userId int64) bool {
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
		close(LocalChan)
		close(SWChan)
		close(CASChan)
		return false
	}
}
