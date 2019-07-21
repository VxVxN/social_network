package language

import (
	"encoding/json"
	"fmt"
	"net/http"
	cnfg "social_network/src/config"
	"social_network/src/log"
)

type response struct {
	Language string `json:language`
}

func GetLanguage(w http.ResponseWriter, r *http.Request) {
	var lang string
	langCookie, err := r.Cookie("language")
	if err != nil {
		lang = cnfg.Config.DefaultLanguage
	} else {
		lang = langCookie.Value
	}

	var resp response
	resp.Language = lang

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.ComLog.Error.Printf("Error marshal response: %v", err)
		return
	}
	fmt.Fprintln(w, string(respJSON))
}
