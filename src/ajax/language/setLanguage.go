package language

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"social_network/src/log"
	resp "social_network/src/response"
)

type requestLanguage struct {
	Language string `json:"language"`
}

func SetLanguage(w http.ResponseWriter, r *http.Request) resp.Response {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.ComLog.Error.Printf("Failed to redding body. Error: %v", err)
		return resp.Error400("Error redding body")
	}
	var req requestLanguage
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.ComLog.Error.Printf("Failed to unmarshal body. Error: %v", err)
		return resp.Error400("Failed to unmarshal body")
	}

	cookie := http.Cookie{Name: "language", Value: req.Language}

	http.SetCookie(w, &cookie)

	return resp.Success(nil)
}
