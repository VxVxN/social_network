package common

import (
	"net/http"
	"time"

	"github.com/VxVxN/social_network/cmd/ajax_server/context"
	"github.com/VxVxN/social_network/internal/tools"
)

type responseListUsers struct {
	Nicknames []string `json:"nicknames"`
}

func ListUsers(w http.ResponseWriter, r *http.Request, ctx *context.Context) tools.Response {
	c, err := r.Cookie("session_token")
	if err != nil {
		ctx.Log.Error.Printf("Error get session token: %v", err)
		return tools.Error400("Failed to get cookie")
	}
	sessionToken := c.Value

	timeAddMinute := time.Now().Add(-time.Minute)
	rows, err := ctx.Database.Query("SELECT nickname FROM users WHERE id IN (SELECT user_id FROM sessions WHERE last_online>? AND session!=?)", timeAddMinute, sessionToken)
	if err != nil {
		ctx.Log.Error.Printf("Error get list users: %v", err)
		return tools.Error500("Failed to get users")
	}
	defer rows.Close()
	listUsers := responseListUsers{Nicknames: []string{}}
	for rows.Next() {
		var nickname string
		if err := rows.Scan(&nickname); err != nil {
			return tools.Error500("Failed to scan users")
		}
		listUsers.Nicknames = append(listUsers.Nicknames, nickname)
	}
	return tools.Success(listUsers)

}
