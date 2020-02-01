package authorization

import (
	"database/sql"
	"html/template"
	"net/http"
	"strings"

	"social_network/cmd/web_server/context"

	"golang.org/x/crypto/bcrypt"
)

type registrationPage struct {
	Page
	IsRegistration bool
}

var rPage = registrationPage{}
var registrationTemplate = template.Must(template.New("main").ParseFiles("web/templates/registration.html"))

func RegistrationForm(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
	rPage.IsRegistration = true
	rPage.Action = "/registration"
	rPage.GoMain = "/"

	registrationTemplate.ExecuteTemplate(w, "registration.html", rPage)
}

func Registration(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
	username := r.FormValue("username")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	nicknameRow := ctx.Database.QueryRow("SELECT nickname FROM users WHERE nickname=?", username)
	emailRow := ctx.Database.QueryRow("SELECT email FROM users WHERE email=?", email)
	errNickname := nicknameRow.Scan()
	errEmail := emailRow.Scan()
	if errEmail != sql.ErrNoRows || errNickname != sql.ErrNoRows {
		rPage.Error = ""
		if errEmail != sql.ErrNoRows {
			rPage.Error = "Email already exists."
		}
		if errNickname != sql.ErrNoRows {
			rPage.Error = "Nickname already exists."
		}
		registrationTemplate.ExecuteTemplate(w, "registration.html", rPage)
	} else {
		password, err := hashAndSalt([]byte(password))
		if err != nil {
			ctx.Log.Error.Println(err)
			return
		}
		email = strings.ToLower(email)
		row := ctx.Database.QueryRow("INSERT INTO users (nickname, fname, lname, email, password) VALUES (?, ?, ?, ?, ?)", username, fname, lname, email, password)
		_ = row.Scan()
		http.Redirect(w, r, "/authorization", http.StatusMovedPermanently)
	}
}

func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
