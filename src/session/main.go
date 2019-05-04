package session

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type Page struct {
	Title    string
	Nickname string
}

type mainPage struct {
	Page
	LogOut string
}

var Database *sql.DB

var aPage = mainPage{}
var mainTemplate = template.Must(template.New("main").ParseFiles("templates/main.html"))

var Store = sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))

func MainPage(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	aPage.Nickname = session.Values["nickname"].(string)
	aPage.Title = "Main"
	aPage.LogOut = "/logout"
	mainTemplate.ExecuteTemplate(w, "main.html", aPage)
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["id"] = -1
	session.Values["nickname"] = ""
	session.Values["authenticated"] = false
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
