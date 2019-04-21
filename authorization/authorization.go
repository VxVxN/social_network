package authorization

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
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
	aPage.GoMain = "/main"
	aPage.Error = ""
	authorizationTemplate.ExecuteTemplate(w, "authorization.html", aPage)
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if r.Method == "POST" {
		uerr := Database.QueryRow("SELECT * FROM users WHERE email=? AND password=?", email, password)
		if uerr.Scan() == sql.ErrNoRows {
			err := errors.New("User is not found")
			fmt.Println(err)
			aPage.Error = err.Error()
			authorizationTemplate.ExecuteTemplate(w, "authorization.html", aPage)
		} else {
			http.Redirect(w, r, "", http.StatusAccepted)
		}
	}
}
