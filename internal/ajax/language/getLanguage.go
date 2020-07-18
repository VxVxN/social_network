package language

import (
	"net/http"
	"social_network/cmd/ajax_server/context"
	cnfg "social_network/internal/config"
	"social_network/internal/tools"
)

// TODO: move it
var allLangs = []string{"EN", "RU"} // const

func GetLanguage(w http.ResponseWriter, r *http.Request, ctx *context.Context) tools.Response {
	var isCookie bool
	lang := cnfg.Config.DefaultLanguage

	langCookie, err := r.Cookie("language")
	if err == nil {
		isCookie = true
		lang = langCookie.Value
	}

	if !tools.ContainsString(lang, allLangs) {
		ctx.Log.Error.Printf("Invalid language. Value: %s. IsCookie: %v", lang, isCookie)
		return tools.Error500("Invalid language")
	}

	return tools.Success(lang)
}
