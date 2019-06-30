package online

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	app "social_network/src/application"
)

type requestOnline struct {
	Online bool `json:online`
}

func SetOnline(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		app.ComLog.Error.Printf("Error redding body: %v", err)
		return
	}
	var req requestOnline
	err = json.Unmarshal(body, &req)
	if err != nil {
		app.ComLog.Error.Printf("Error redding body: %v", err)
		return
	}

	c, err := r.Cookie("session_token")
	if err != nil {
		app.ComLog.Error.Printf("Error get session_token: %v", err)
		return
	}
	sessionToken := c.Value

	_ = app.Database.QueryRow("UPDATE sessions SET last_online=? WHERE session=?", time.Now(), sessionToken)
}
