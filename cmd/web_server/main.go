package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	"social_network/cmd/web_server/context"
	"social_network/internal/authorization"
	cnfg "social_network/internal/config"
	"social_network/internal/log"
	"social_network/internal/session"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type Page struct {
	Title  string
	LogIn  string
	SignUp string
}

func main() {
	context := &context.Context{Log: log.Init("web_server.log")}

	context.Log.Info.Println("Web server start.")
	mysqlPort := strconv.Itoa(cnfg.Config.MysqlPort)
	db, err := sql.Open("mysql", cnfg.Config.MysqlName+":"+cnfg.Config.MysqlPassword+"@tcp("+cnfg.Config.MysqlIP+":"+mysqlPort+")/social_network")
	if err != nil {
		context.Log.Fatal.Printf("Error open mysql: %v", err)
		panic(err)
	}
	defer db.Close()

	context.Database = db

	routes := httprouter.New()
	routes.GET("/", middleware(mainForm, context))
	routes.GET("/main", middleware(session.MainPage, context))
	routes.GET("/logout", middleware(session.LogOut, context))
	routes.POST("/authorization", middleware(authorization.Authorize, context))
	routes.GET("/authorization", middleware(authorization.AuthorizationForm, context))
	routes.POST("/registration", middleware(authorization.Registration, context))
	routes.GET("/registration", middleware(authorization.RegistrationForm, context))

	fs := http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", routes)
	port := ":" + strconv.Itoa(cnfg.Config.WebServerPort)
	if err := http.ListenAndServe(port, nil); err != nil {
		context.Log.Fatal.Printf("Failed to listen and serve port: %v. Error: %v", port, err)
		panic(err)
	}
}

func mainForm(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
	main := Page{}
	main.Title = "Main"
	main.LogIn = "/authorization"
	main.SignUp = "/registration"
	tpl := template.Must(template.New("main").ParseFiles("web/templates/index.html"))
	tpl.ExecuteTemplate(w, "index.html", main)
}

func middleware(handler func(http.ResponseWriter, *http.Request, *context.Context), ctx *context.Context) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handler(w, r, ctx)
	})
}
