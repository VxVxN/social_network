package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/VxVxN/social_network/app/ajax/common"
	"github.com/VxVxN/social_network/app/ajax/language"
	"github.com/VxVxN/social_network/app/ajax/online"
	cnfg "github.com/VxVxN/social_network/app/config"
	"github.com/VxVxN/social_network/app/log"
	"github.com/VxVxN/social_network/app/tools"
	"github.com/VxVxN/social_network/cmd/ajax_server/context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	context := &context.Context{Log: log.Init("ajax_server.log", false)}

	context.Log.Info.Println("Ajax server start.")
	mysqlPort := strconv.Itoa(cnfg.Config.MysqlPort)
	db, err := sql.Open("mysql", cnfg.Config.MysqlName+":"+cnfg.Config.MysqlPassword+"@tcp("+cnfg.Config.MysqlIP+":"+mysqlPort+")/social_network")
	if err != nil {
		context.Log.Fatal.Printf("Error open mysql: %v", err)
		panic(err)
	}
	defer db.Close()

	context.Database = db

	routes := httprouter.New()
	routes.POST("/ajax/language", middleware(language.SetLanguage, context))
	routes.GET("/ajax/language", middleware(language.GetLanguage, context))

	routes.POST("/ajax/online", middleware(online.SetOnline, context))

	routes.GET("/ajax/list_users", middleware(common.ListUsers, context))
	routes.GET("/ajax/get_messages", middleware(common.GetMessages, context))
	routes.POST("/ajax/send_message", middleware(common.SendMessage, context))

	http.Handle("/", routes)
	port := ":" + strconv.Itoa(cnfg.Config.AJAXServerPort)
	if err := http.ListenAndServe(port, nil); err != nil {
		context.Log.Fatal.Printf("Failed to listen and serve port: %v. Error: %v", port, err)
		panic(err)
	}
}

func middleware(handler func(http.ResponseWriter, *http.Request, *context.Context) tools.Response, context *context.Context) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		resp := handler(w, r, context)
		respBytes, err := json.Marshal(resp)
		if err != nil {
			context.Log.Error.Printf("Failed to marshal response in middlewareResponse. Error: %v", err)
		}
		w.Write(respBytes)
	})
}
