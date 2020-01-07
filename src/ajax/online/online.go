package online

import (
	"net/http"
	"time"

	app "social_network/src/application"
	"social_network/src/log"
	"social_network/src/tools"
)

type requestOnline struct {
	Online bool `json:online`
}

func SetOnline(w http.ResponseWriter, r *http.Request) tools.Response {

	var req requestOnline
	if err := tools.UnmarshalRequest(r.Body, &req); err != nil {
		log.ComLog.Error.Printf("Failed to unmarshal body. Error: %v", err)
		return tools.Error400("Failed to unmarshal body")
	}

	c, err := r.Cookie("session_token")
	if err != nil {
		log.ComLog.Error.Printf("Error get session_token: %v", err)
		return tools.Error400("Failed to get cookie")
	}
	sessionToken := c.Value

	row := app.Database.QueryRow("UPDATE sessions SET last_online=? WHERE session=?", time.Now(), sessionToken)
	_ = row.Scan()
	return tools.Success(nil)
}
