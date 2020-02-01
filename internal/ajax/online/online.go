package online

import (
	"net/http"
	"time"

	"social_network/cmd/ajax_server/context"
	"social_network/internal/tools"
)

type requestOnline struct {
	Online bool `json:online`
}

func SetOnline(w http.ResponseWriter, r *http.Request, ctx *context.Context) tools.Response {

	var req requestOnline
	if err := tools.UnmarshalRequest(r.Body, &req); err != nil {
		ctx.Log.Error.Printf("Failed to unmarshal body. Error: %v", err)
		return tools.Error400("Failed to unmarshal body")
	}

	c, err := r.Cookie("session_token")
	if err != nil {
		ctx.Log.Error.Printf("Error get session_token: %v", err)
		return tools.Error400("Failed to get cookie")
	}
	sessionToken := c.Value

	row := ctx.Database.QueryRow("UPDATE sessions SET last_online=? WHERE session=?", time.Now(), sessionToken)
	_ = row.Scan()
	return tools.Success(nil)
}
