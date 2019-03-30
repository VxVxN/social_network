package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"social_network/authorization"

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
	routes.HandleFunc("/main", mainForm).Methods("GET")
	routes.HandleFunc("/authorization", authorization.Authorize).Methods("POST")
	routes.HandleFunc("/authorization", authorization.AuthorizationForm).Methods("GET")
	http.Handle("/", routes)
	http.ListenAndServe(":8080", nil)
}

func mainForm(w http.ResponseWriter, r *http.Request) {
	main := Page{}
	main.Title = "Main"
	main.LogIn = "/authorization"
	main.SignUp = "/main"
	tpl := template.Must(template.New("main").ParseFiles("ui/html/main.html"))
	tpl.ExecuteTemplate(w, "main.html", main)
}
