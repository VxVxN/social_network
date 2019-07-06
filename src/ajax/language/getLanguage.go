package language

import (
	"encoding/json"
	"fmt"
	"net/http"
	app "social_network/src/application"
)

type response struct {
	Language string `json:language`
}

func GetLanguage(w http.ResponseWriter, r *http.Request) {
	var lang string
	langCookie, err := r.Cookie("language")
	if err != nil {
		lang = "RU"
	} else {
		lang = langCookie.Value
	}

	var resp response
	resp.Language = lang

	respJSON, err := json.Marshal(resp)
	if err != nil {
		app.ComLog.Error.Printf("Error marshal response: %v", err)
		return
	}
	fmt.Fprintln(w, string(respJSON))
}
