package online

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	app "social_network/src/application"
	"social_network/src/log"
	resp "social_network/src/response"
)

type requestOnline struct {
	Online bool `json:online`
}

func SetOnline(w http.ResponseWriter, r *http.Request) resp.Response {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.ComLog.Error.Printf("Error redding body: %v", err)
		return resp.Error400("Error redding body")
	}
	var req requestOnline
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.ComLog.Error.Printf("Error redding body: %v", err)
		return resp.Error400("Failed to unmarshal body")
	}

	c, err := r.Cookie("session_token")
	if err != nil {
		log.ComLog.Error.Printf("Error get session_token: %v", err)
		return resp.Error400("Failed to get cookie")
	}
	sessionToken := c.Value

	row := app.Database.QueryRow("UPDATE sessions SET last_online=? WHERE session=?", time.Now(), sessionToken)
	_ = row.Scan()
	return resp.Success(nil)
}
