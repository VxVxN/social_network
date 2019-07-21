package language

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"social_network/src/log"
)

type requestLanguage struct {
	Language string `json:language`
}

func SetLanguage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.ComLog.Error.Printf("Error redding body: %v", err)
		return
	}
	var req requestLanguage
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.ComLog.Error.Printf("Error redding body: %v", err)
		return
	}

	cookie := http.Cookie{Name: "language", Value: req.Language}

	http.SetCookie(w, &cookie)
}
