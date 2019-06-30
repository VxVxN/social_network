package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	app "social_network/src/application"
)

type response struct {
	Nickname []string `json:nickname`
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	timeAddMinute := time.Now().Add(-time.Minute)
	rows, err := app.Database.Query("SELECT nickname FROM users WHERE id IN (SELECT user_id FROM sessions WHERE last_online>?)", timeAddMinute)
	if err != nil {
		app.ComLog.Error.Printf("Error get list users: %v", err)
		return
	}

	resp := response{}
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
