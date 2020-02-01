package language

import (
	"net/http"
	"social_network/cmd/ajax_server/context"
	"social_network/internal/tools"
)

type requestLanguage struct {
	Language string `json:"language"`
}

func SetLanguage(w http.ResponseWriter, r *http.Request, ctx *context.Context) tools.Response {

	var req requestLanguage
	if err := tools.UnmarshalRequest(r.Body, &req); err != nil {
		ctx.Log.Error.Printf("Failed to unmarshal body. Error: %v", err)
		return tools.Error400("Failed to unmarshal body")
	}

	cookie := http.Cookie{Name: "language", Value: req.Language}

	http.SetCookie(w, &cookie)

	return tools.Success(nil)
}