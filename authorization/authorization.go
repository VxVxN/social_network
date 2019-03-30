package authorization

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title          string
	Action         string
	Application    string
	IsAuthenticate bool
	isPermission   bool
	GoMain         string
}

var Database *sql.DB

func AuthorizationForm(w http.ResponseWriter, r *http.Request) {
	Authorize := Page{}
	Authorize.Title = "Login"
	Authorize.IsAuthenticate = true
	Authorize.Application = ""
	Authorize.Action = "/authorization"
	Authorize.GoMain = "/main"
	tpl := template.Must(template.New("main").ParseFiles("ui/html/authorization.html"))
	tpl.ExecuteTemplate(w, "authorization.html", Authorize)
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	isPermission := r.FormValue("isPermission")

	uerr := Database.QueryRow("SELECT * from users where nickname=? and password=?", username, password)
	if uerr != nil {
		err := errors.New("User is not found.")
		fmt.Println(err)
		fmt.Fprintln(w, err.Error())
		return
	}

	if isPermission == "1" {
		http.Redirect(w, r, "", http.StatusAccepted)
	} else {
		http.Redirect(w, r, "/authorization", http.StatusUnauthorized)
	}
}
