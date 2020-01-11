package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"social_network/internal/ajax/common"
	"social_network/internal/ajax/language"
	"social_network/internal/ajax/online"
	app "social_network/internal/application"
	cnfg "social_network/internal/config"
	"social_network/internal/log"
	"social_network/internal/tools"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log.ComLog.Info.Println("Ajax server start.")
	mysqlPort := strconv.Itoa(cnfg.Config.MysqlPort)
	db, err := sql.Open("mysql", cnfg.Config.MysqlName+":"+cnfg.Config.MysqlPassword+"@tcp("+cnfg.Config.MysqlIP+":"+mysqlPort+")/social_network")
	if err != nil {
		log.ComLog.Fatal.Printf("Error open mysql: %v", err)
		panic(err)
	}
	defer db.Close()

	app.Database = db

	routes := httprouter.New()
	routes.POST("/ajax/language", middlewareResponse(language.SetLanguage))
	routes.GET("/ajax/language", middlewareResponse(language.GetLanguage))

	routes.POST("/ajax/online", middlewareResponse(online.SetOnline))

	routes.GET("/ajax/list_users", middlewareResponse(common.ListUsers))
	routes.GET("/ajax/get_messages", middlewareResponse(common.GetMessages))
	routes.POST("/ajax/send_message", middlewareResponse(common.SendMessage))

	http.Handle("/", routes)
	port := ":" + strconv.Itoa(cnfg.Config.AJAXServerPort)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.ComLog.Fatal.Printf("Failed to listen and serve port: %v. Error: %v", port, err)
		panic(err)
	}
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
