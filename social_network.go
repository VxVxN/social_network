package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"social_network/src/authorization"
	"social_network/src/session"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Page struct {
	Title  string
	LogIn  string
	SignUp string
}

func main() {
	fmt.Println("Server start.")
	db, err := sql.Open("mysql", "user:123@tcp(127.0.0.1:3306)/social_network")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	authorization.Database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/", mainForm).Methods("GET")
	routes.HandleFunc("/main", session.MainPage).Methods("GET")
	routes.HandleFunc("/logout", session.LogOut).Methods("GET")
	routes.HandleFunc("/authorization", authorization.Authorize).Methods("POST")
	routes.HandleFunc("/authorization", authorization.AuthorizationForm).Methods("GET")
	routes.HandleFunc("/registration", authorization.Registration).Methods("POST")
	routes.HandleFunc("/registration", authorization.RegistrationForm).Methods("GET")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", routes)
	http.ListenAndServe(":8080", nil)
}

func mainForm(w http.ResponseWriter, r *http.Request) {
	main := Page{}
	main.Title = "Main"
	main.LogIn = "/authorization"
	main.SignUp = "/registration"
	tpl := template.Must(template.New("main").ParseFiles("templates/index.html"))
	tpl.ExecuteTemplate(w, "index.html", main)
}
