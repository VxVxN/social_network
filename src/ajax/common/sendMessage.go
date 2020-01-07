package common

import (
	"net/http"
	app "social_network/src/application"
	"social_network/src/log"
	"social_network/src/tools"
	"time"
)

type requestSendMessage struct {
	Nickname string `json:nickname`
	Message  string `json:message`
}

func SendMessage(w http.ResponseWriter, r *http.Request) tools.Response {

	var req requestSendMessage
	if err := tools.UnmarshalRequest(r.Body, &req); err != nil {
		log.ComLog.Error.Printf("Failed to unmarshal body. Error: %v", err)
		return tools.Error400("Failed to unmarshal body")
	}

	c, err := r.Cookie("session_token")
	if err != nil {
		log.ComLog.Error.Printf("Error get session token: %v", err)
		return tools.Error400("Failed to get cookie")
	}
	sessionToken := c.Value

	row := app.Database.QueryRow("SELECT id FROM users WHERE nickname=?", req.Nickname)
	var secondID int
	if err = row.Scan(&secondID); err != nil {
		log.ComLog.Error.Printf("Error get id by nickname: %v. Error: %v", req.Nickname, err)
		return tools.Error500("Failed to get id by nickname")
	}

	row = app.Database.QueryRow("SELECT user_id FROM sessions WHERE session=?", sessionToken)
	var firstID int
	if err = row.Scan(&firstID); err != nil {
		log.ComLog.Error.Printf("Error get id user: %v", err)
		return tools.Error500("Failed to get id user")
	}
	row = app.Database.QueryRow("INSERT INTO messages (first_id, message, second_id, time_sending) VALUES (?, ?, ?, ?)", firstID, req.Message, secondID, time.Now())
	_ = row.Scan()

	return tools.Success(nil)
}
