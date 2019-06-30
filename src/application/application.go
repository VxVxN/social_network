package application

import (
	"database/sql"
	log "social_network/src/logger"
)

var (
	Database *sql.DB

	ComLog *log.Logger
)

func init() {
	ComLog = log.Init("common.log")
}
