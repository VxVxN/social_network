package session

import (
	"database/sql"
	"html/template"
	"net/http"

	"social_network/cmd/web_server/context"
	app "social_network/internal/application"

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
var mainTemplate = template.Must(template.New("main").ParseFiles("web/templates/main.html"))

var Store = sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))

func MainPage(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	row := app.Database.QueryRow("SELECT nickname FROM users WHERE id IN (SELECT user_id FROM sessions WHERE session=?)", sessionToken)
	var nickname string
	err = row.Scan(&nickname)
	if err != nil {
		ctx.Log.Error.Printf("Error get user: %v", err)
		return
	}
	aPage.Nickname = nickname
	aPage.Title = "Main"
	aPage.LogOut = "/logout"
	mainTemplate.ExecuteTemplate(w, "main.html", aPage)
}

func LogOut(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
	w.Header().Set("Cache-Control", "no-cache")
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value
	row := app.Database.QueryRow("DELETE FROM sessions WHERE session=?", sessionToken)
	_ = row.Scan()

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
