package authorization

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"social_network/src/session"
)

type Page struct {
	Title  string
	Action string
	GoMain string
	Error  string
}

type authorizationPage struct {
	Page
	IsAuthenticate bool
}

var Database *sql.DB

var aPage = authorizationPage{}
var authorizationTemplate = template.Must(template.New("main").ParseFiles("templates/authorization.html"))

func AuthorizationForm(w http.ResponseWriter, r *http.Request) {
	aPage.Title = "Login"
	aPage.IsAuthenticate = true
	aPage.Action = "/authorization"
	aPage.GoMain = "/"
	aPage.Error = ""
	authorizationTemplate.ExecuteTemplate(w, "authorization.html", aPage)
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	row := Database.QueryRow("SELECT id, nickname FROM users WHERE email=? AND password=?", email, password)
	var id int
	var nickname string
	err := row.Scan(&id, &nickname)
	if err != nil {
		err := errors.New("User is not found")
		fmt.Println(err)
		aPage.Error = err.Error()
		authorizationTemplate.ExecuteTemplate(w, "authorization.html", aPage)
	} else {
		session, err := session.Store.Get(r, "session")
		if err == nil {
			session.Values["id"] = id
			session.Values["nickname"] = nickname
			session.Values["authenticated"] = true
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/main", http.StatusMovedPermanently)
		}
	}
}
