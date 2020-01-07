package common

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	app "social_network/src/application"
	"social_network/src/log"
	"social_network/src/tools"
)

type responseMessages struct {
	Nickname string    `json:"nickname"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time_sending"`
}

func GetMessages(w http.ResponseWriter, r *http.Request) tools.Response {
	secondNickname := r.FormValue("nickname")

	c, err := r.Cookie("session_token")
	if err != nil {
		log.ComLog.Error.Printf("Error get session token: %v", err)
		return tools.Error400("Failed to get cookie")
	}
	sessionToken := c.Value

	row := app.Database.QueryRow("SELECT user_id FROM sessions WHERE session=?", sessionToken)
	var firstID int
	if err = row.Scan(&firstID); err != nil {
		log.ComLog.Error.Printf("Error get id user: %v", err)
		return tools.Error500("Failed to get first user id")
	}

	row = app.Database.QueryRow("SELECT nickname FROM users WHERE id=?", firstID)
	var firstNickname string
	err = row.Scan(&firstNickname)
	if err != nil {
		log.ComLog.Error.Printf("Error get user: %v", err)
		return tools.Error500("Failed to get nickname")
	}

	row = app.Database.QueryRow("SELECT id FROM users WHERE nickname=?", secondNickname)
	var secondID int
	err = row.Scan(&secondID)
	if err != nil {
		log.ComLog.Error.Printf("Error get id by nickname: %v. Error: %v", secondNickname, err)
		return tools.Error500("Failed to get second user id")
	}

	var messagesResult []responseMessages

	messagesResult, err = getMessages(firstNickname, firstID, secondID)
	if err != nil {
		log.ComLog.Error.Printf("Error get messages: %v", err)
		return tools.Error500("Failed to get messages")
	}
	secondMessages, err := getMessages(secondNickname, secondID, firstID)
	if err != nil {
		log.ComLog.Error.Printf("Error get messages: %v", err)
		return tools.Error500("Failed to get messages")
	}

	messagesResult = append(messagesResult, secondMessages...)
	sort.Slice(messagesResult, func(i, j int) bool { return messagesResult[i].Time.Before(messagesResult[j].Time) })

	return tools.Success(messagesResult)
}

func getMessages(nickname string, firstID, secondID int) ([]responseMessages, error) {
	rows, err := app.Database.Query("SELECT message, time_sending FROM messages WHERE first_id=? AND second_id=?", firstID, secondID)
	if err != nil {
		errText := fmt.Sprintf("Error get list messages: %v", err)
		return nil, errors.New(errText)
	}
	defer rows.Close()
	messagesResult := []responseMessages{}
	var message responseMessages
	var timeSending string
	for rows.Next() {
		if err := rows.Scan(&message.Message, &timeSending); err != nil {
			return nil, err
		}
		timeSending += "Z"
		timeSending := strings.Replace(timeSending, " ", "T", 1)
		message.Time, _ = time.Parse(time.RFC3339, timeSending)
		message.Nickname = nickname
		messagesResult = append(messagesResult, message)
	}
	return messagesResult, nil
}
