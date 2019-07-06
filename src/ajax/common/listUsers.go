package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	app "social_network/src/application"
)

type responseListUsers struct {
	Nickname []string `json:nickname`
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		app.ComLog.Error.Printf("Error get session token: %v", err)
		return
	}
	sessionToken := c.Value

	timeAddMinute := time.Now().Add(-time.Minute)
	rows, err := app.Database.Query("SELECT nickname FROM users WHERE id IN (SELECT user_id FROM sessions WHERE last_online>? AND session!=?)", timeAddMinute, sessionToken)
	if err != nil {
		app.ComLog.Error.Printf("Error get list users: %v", err)
		return
	}

	resp := responseListUsers{}
	for rows.Next() {
		var nickname string
		rows.Scan(&nickname)
		resp.Nickname = append(resp.Nickname, nickname)
	}
	output, err := json.Marshal(resp)
	if err != nil {
		app.ComLog.Error.Printf("Error marshal response: %v", err)
		return
	}
	fmt.Fprintln(w, string(output))
}
