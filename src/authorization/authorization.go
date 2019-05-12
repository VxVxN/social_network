package authorization

import (
	"errors"
	"html/template"
	"net/http"

	app "social_network/src/application"

	uuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	row := app.Database.QueryRow("SELECT id, nickname, password FROM users WHERE email=?", email)
	var id int
	var nickname, hashPassword string
	err := row.Scan(&id, &nickname, &hashPassword)
	if err != nil {
		err := errors.New("User is not found")
		aPage.Error = err.Error()
		authorizationTemplate.ExecuteTemplate(w, "authorization.html", aPage)
	} else {
		if comparePasswords(hashPassword, []byte(password)) {
			sessionToken := uuid.New().String()
			_ = app.Database.QueryRow("INSERT INTO sessions (session, user_id) VALUES (?, ?)", sessionToken, id)

			http.SetCookie(w, &http.Cookie{
				Name:  "session_token",
				Value: sessionToken,
			})
			http.Redirect(w, r, "/main", http.StatusMovedPermanently)
		} else {
			err := errors.New("User is not found")
			aPage.Error = err.Error()
			authorizationTemplate.ExecuteTemplate(w, "authorization.html", aPage)
		}
	}
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
