package language

import (
	"net/http"
	cnfg "social_network/src/config"
	"social_network/src/log"
	resp "social_network/src/response"
)

func GetLanguage(w http.ResponseWriter, r *http.Request) resp.Response {
	var lang string
	langCookie, err := r.Cookie("language")
	if err != nil {
		lang = cnfg.Config.DefaultLanguage
	} else {
		lang = langCookie.Value
	}
	if lang == "" {
		log.ComLog.Error.Printf("Failed to get language")
		return resp.Error500("Error to get language")
	}
	return resp.Success(lang)
}
