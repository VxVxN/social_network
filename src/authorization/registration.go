package authorization

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type registrationPage struct {
	Page
	IsRegistration bool
	Username       string
	Fname          string
	Lname          string
	Email          string
	Password       string
}

var rPage = registrationPage{}
var registrationTemplate = template.Must(template.New("main").ParseFiles("templates/registration.html"))

func RegistrationForm(w http.ResponseWriter, r *http.Request) {
	rPage.Title = "Sigh up"
	rPage.IsRegistration = true
	rPage.Action = "/registration"
	rPage.GoMain = "/"
	rPage.Username = ""
	rPage.Fname = ""
	rPage.Lname = ""
	rPage.Email = ""
	rPage.Password = ""

	registrationTemplate.ExecuteTemplate(w, "registration.html", rPage)
}

func Registration(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	nicknameRow := Database.QueryRow("SELECT nickname FROM users WHERE nickname=?", username)
	emailRow := Database.QueryRow("SELECT email FROM users WHERE email=?", email)
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
		rPage.Username = username
		rPage.Fname = fname
		rPage.Lname = lname
		rPage.Email = email
		rPage.Password = password
		registrationTemplate.ExecuteTemplate(w, "registration.html", rPage)
	} else {
		password = hashAndSalt([]byte(password))
		email = strings.ToLower(email)
		_ = Database.QueryRow("INSERT INTO users (nickname, fname, lname, email, password) VALUES (?, ?, ?, ?, ?)", username, fname, lname, email, password)
		http.Redirect(w, r, "/authorization", http.StatusMovedPermanently)
	}
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
