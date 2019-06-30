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
	lang, err := r.Cookie("language")
	if err != nil {
		w.Write([]byte("EN"))
	}

	var resp response
	resp.Language = lang.Value

	respJSON, err := json.Marshal(resp)
	if err != nil {
		app.ComLog.Error.Printf("Error marshal response: %v", err)
		return
	}
	fmt.Fprintln(w, string(respJSON))
}
