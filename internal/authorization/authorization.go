package authorization

import (
	"errors"
	"html/template"
	"net/http"
	"time"

	"social_network/cmd/web_server/context"

	uuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Page struct {
	Action string
	GoMain string
	Error  string
}

type authorizationPage struct {
	Page
	IsAuthenticate bool
}

var aPage = authorizationPage{}
var authorizationTemplate = template.Must(template.New("main").ParseFiles("web/templates/authorization.html"))

func AuthorizationForm(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
	w.Header().Set("Cache-Control", "no-cache")

	c, err := r.Cookie("session_token")
	if err == nil {
		sessionToken := c.Value
		row := ctx.Database.QueryRow("SELECT user_id FROM sessions WHERE session=?", sessionToken)
		var userID int
		err = row.Scan(&userID)
		if err == nil && userID != 0 {
			http.Redirect(w, r, "/main", http.StatusMovedPermanently)
		}
	}
	aPage.IsAuthenticate = true
	aPage.Action = "/authorization"
	aPage.GoMain = "/"
	aPage.Error = ""
	authorizationTemplate.ExecuteTemplate(w, "authorization.html", aPage)
}

func Authorize(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	row := ctx.Database.QueryRow("SELECT id, nickname, password FROM users WHERE email=?", email)
	var id int
	var nickname, hashPassword string
	err := row.Scan(&id, &nickname, &hashPassword)
	if err != nil {
		ctx.Log.Error.Printf("Error geting user data: %v", err)
		aPage.Error = errors.New("User is not found").Error()
		authorizationTemplate.ExecuteTemplate(w, "authorization.html", aPage)
	} else {
		if comparePasswords(hashPassword, []byte(password)) {
			sessionToken := uuid.New().String()
			row = ctx.Database.QueryRow("INSERT INTO sessions (session, user_id, last_online) VALUES (?, ?, ?)", sessionToken, id, time.Now())

			err := row.Scan()
			if err != nil {
				ctx.Log.Error.Printf("Error create session: %v", err)
				return
			}

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
