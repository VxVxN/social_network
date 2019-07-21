package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"social_network/src/log"
)

type config struct {
	Port            int    `json:"PORT"`
	MysqlName       string `json:"MYSQL_NAME"`
	MysqlPassword   string `json:"MYSQL_PASSWORD"`
	MysqlIP         string `json:"MYSQL_IP"`
	MysqlPort       int    `json:"MYSQL_PORT"`
	DefaultLanguage string `json:"DEFAULT_LANGUAGE"`
}

var Config config

func init() {
	jsonFile, err := os.Open("config/main.json")
	if err != nil {
		log.ComLog.Fatal.Printf("Error open main config: %v", err)
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.ComLog.Fatal.Printf("Error read main config: %v", err)
		panic(err)
	}
	json.Unmarshal(byteValue, &Config)
}
