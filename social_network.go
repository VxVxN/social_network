package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	"social_network/src/ajax/common"
	"social_network/src/ajax/language"
	"social_network/src/ajax/online"
	app "social_network/src/application"
	"social_network/src/authorization"
	cnfg "social_network/src/config"
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
	app.ComLog.Info.Println("Server start.")
	mysqlPort := strconv.Itoa(cnfg.Config.MysqlPort)
	db, err := sql.Open("mysql", cnfg.Config.MysqlName+":"+cnfg.Config.MysqlPassword+"@tcp("+cnfg.Config.MysqlIP+":"+mysqlPort+")/social_network")
	if err != nil {
		app.ComLog.Fatal.Printf("Error open mysql: %v", err)
		panic(err)
	}
	defer db.Close()

	app.Database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/", mainForm).Methods("GET")
	routes.HandleFunc("/main", session.MainPage).Methods("GET")
	routes.HandleFunc("/logout", session.LogOut).Methods("GET")
	routes.HandleFunc("/authorization", authorization.Authorize).Methods("POST")
	routes.HandleFunc("/authorization", authorization.AuthorizationForm).Methods("GET")
	routes.HandleFunc("/registration", authorization.Registration).Methods("POST")
	routes.HandleFunc("/registration", authorization.RegistrationForm).Methods("GET")

	routes.HandleFunc("/ajax/language", language.SetLanguage).Methods("POST")
	routes.HandleFunc("/ajax/language", language.GetLanguage).Methods("GET")

	routes.HandleFunc("/ajax/online", online.SetOnline).Methods("POST")

	routes.HandleFunc("/ajax/list_users", common.ListUsers).Methods("GET")
	routes.HandleFunc("/ajax/get_messages", common.GetMessages).Methods("GET")
	routes.HandleFunc("/ajax/send_message", common.SendMessage).Methods("POST")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", routes)
	port := ":" + strconv.Itoa(cnfg.Config.Port)
	http.ListenAndServe(port, nil)
}

func mainForm(w http.ResponseWriter, r *http.Request) {
	main := Page{}
	main.Title = "Main"
	main.LogIn = "/authorization"
	main.SignUp = "/registration"
	tpl := template.Must(template.New("main").ParseFiles("templates/index.html"))
	tpl.ExecuteTemplate(w, "index.html", main)
}
