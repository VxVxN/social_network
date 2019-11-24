package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	app "social_network/src/application"
	"social_network/src/log"
	"time"
)

type requestSendMessage struct {
	Nickname string `json:nickname`
	Message  string `json:message`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.ComLog.Error.Printf("Error redding body: %v", err)
		return
	}
	var req requestSendMessage
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.ComLog.Error.Printf("Error redding body: %v", err)
		return
	}

	c, err := r.Cookie("session_token")
	if err != nil {
		log.ComLog.Error.Printf("Error get session token: %v", err)
		return
	}
	sessionToken := c.Value

	row := app.Database.QueryRow("SELECT id FROM users WHERE nickname=?", req.Nickname)
	var secondID int
	err = row.Scan(&secondID)
	if err != nil {
		log.ComLog.Error.Printf("Error get id by nickname: %v. Error: %v", req.Nickname, err)
		return
	}

	row = app.Database.QueryRow("SELECT user_id FROM sessions WHERE session=?", sessionToken)
	var firstID int
	err = row.Scan(&firstID)
	if err != nil {
		log.ComLog.Error.Printf("Error get id user: %v", err)
		return
	}
	row = app.Database.QueryRow("INSERT INTO messages (first_id, message, second_id, time_sending) VALUES (?, ?, ?, ?)", firstID, req.Message, secondID, time.Now())
	_ = row.Scan()
}
