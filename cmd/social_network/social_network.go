package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"social_network/internal/ajax/common"
	"social_network/internal/ajax/language"
	"social_network/internal/ajax/online"
	app "social_network/internal/application"
	"social_network/internal/authorization"
	cnfg "social_network/internal/config"
	"social_network/internal/log"
	"social_network/internal/session"
	"social_network/internal/tools"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type Page struct {
	Title  string
	LogIn  string
	SignUp string
}

func main() {
	log.ComLog.Info.Println("Server start.")
	mysqlPort := strconv.Itoa(cnfg.Config.MysqlPort)
	db, err := sql.Open("mysql", cnfg.Config.MysqlName+":"+cnfg.Config.MysqlPassword+"@tcp("+cnfg.Config.MysqlIP+":"+mysqlPort+")/social_network")
	if err != nil {
		log.ComLog.Fatal.Printf("Error open mysql: %v", err)
		panic(err)
	}
	defer db.Close()

	app.Database = db

	routes := httprouter.New()
	routes.GET("/", mainForm)
	routes.GET("/main", session.MainPage)
	routes.GET("/logout", session.LogOut)
	routes.POST("/authorization", authorization.Authorize)
	routes.GET("/authorization", authorization.AuthorizationForm)
	routes.POST("/registration", authorization.Registration)
	routes.GET("/registration", authorization.RegistrationForm)

	routes.POST("/ajax/language", middlewareResponse(language.SetLanguage))
	routes.GET("/ajax/language", middlewareResponse(language.GetLanguage))

	routes.POST("/ajax/online", middlewareResponse(online.SetOnline))

	routes.GET("/ajax/list_users", middlewareResponse(common.ListUsers))
	routes.GET("/ajax/get_messages", middlewareResponse(common.GetMessages))
	routes.POST("/ajax/send_message", middlewareResponse(common.SendMessage))

	fs := http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", routes)
	port := ":" + strconv.Itoa(cnfg.Config.Port)
	http.ListenAndServe(port, nil)
}

func mainForm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	main := Page{}
	main.Title = "Main"
	main.LogIn = "/authorization"
	main.SignUp = "/registration"
	tpl := template.Must(template.New("main").ParseFiles("web/templates/index.html"))
	tpl.ExecuteTemplate(w, "index.html", main)
}

func middlewareResponse(handler func(w http.ResponseWriter, r *http.Request) tools.Response) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		resp := handler(w, r)
		respBytes, err := json.Marshal(resp)
		if err != nil {
			log.ComLog.Error.Printf("Failed to marshal response in middlewareResponse. Error: %v", err)
		}
		w.Write(respBytes)
	})
}
