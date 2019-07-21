package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	app "social_network/src/application"
	"social_network/src/log"
)

type responseMessages struct {
	Nickname string    `json:nickname`
	Message  string    `json:message`
	Time     time.Time `json:time_sending`
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	secondNickname := r.FormValue("nickname")

	c, err := r.Cookie("session_token")
	if err != nil {
		log.ComLog.Error.Printf("Error get session token: %v", err)
		return
	}
	sessionToken := c.Value

	row := app.Database.QueryRow("SELECT user_id FROM sessions WHERE session=?", sessionToken)
	var firstID int
	err = row.Scan(&firstID)
	if err != nil {
		log.ComLog.Error.Printf("Error get id user: %v", err)
		return
	}

	row = app.Database.QueryRow("SELECT nickname FROM users WHERE id=?", firstID)
	var firstNickname string
	err = row.Scan(&firstNickname)
	if err != nil {
		log.ComLog.Error.Printf("Error get user: %v", err)
		return
	}

	row = app.Database.QueryRow("SELECT id FROM users WHERE nickname=?", secondNickname)
	var secondID int
	err = row.Scan(&secondID)
	if err != nil {
		log.ComLog.Error.Printf("Error get id by nickname: %v. Error: %v", secondNickname, err)
		return
	}

	var messagesResult []responseMessages

	messagesResult, err = getMessages(firstNickname, firstID, secondID)
	if err != nil {
		log.ComLog.Error.Printf("Error get messages: %v", err)
		return
	}
	secondMessages, err := getMessages(secondNickname, secondID, firstID)
	if err != nil {
		log.ComLog.Error.Printf("Error get messages: %v", err)
		return
	}

	messagesResult = append(messagesResult, secondMessages...)
	sort.Slice(messagesResult, func(i, j int) bool { return messagesResult[i].Time.Before(messagesResult[j].Time) })

	output, err := json.Marshal(messagesResult)
	if err != nil {
		log.ComLog.Error.Printf("Error marshal response: %v", err)
		return
	}
	fmt.Fprintln(w, string(output))
}

func getMessages(nickname string, firstID, secondID int) ([]responseMessages, error) {
	rows, err := app.Database.Query("SELECT message, time_sending FROM messages WHERE first_id=? AND second_id=?", firstID, secondID)
	if err != nil {
		errText := fmt.Sprintf("Error get list messages: %v", err)
		return nil, errors.New(errText)
	}

	var messagesResult []responseMessages
	var message responseMessages
	var timeSending string
	for rows.Next() {
		rows.Scan(&message.Message, &timeSending)
		timeSending += "Z"
		timeSending := strings.Replace(timeSending, " ", "T", 1)
		message.Time, _ = time.Parse(time.RFC3339, timeSending)
		message.Nickname = nickname
		messagesResult = append(messagesResult, message)
	}
	return messagesResult, nil
}
