package language

import (
	"net/http"
	cnfg "social_network/internal/config"
	"social_network/internal/log"
	"social_network/internal/tools"
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
