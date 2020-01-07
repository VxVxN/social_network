package language

import (
	"net/http"
	cnfg "social_network/src/config"
	"social_network/src/log"
	"social_network/src/tools"
)

func GetLanguage(w http.ResponseWriter, r *http.Request) tools.Response {
	var lang string
	langCookie, err := r.Cookie("language")
	if err != nil {
		lang = cnfg.Config.DefaultLanguage
	} else {
		lang = langCookie.Value
	}
	if lang == "" {
		log.ComLog.Error.Printf("Failed to get language")
		return tools.Error500("Error to get language")
	}
	return tools.Success(lang)
}
