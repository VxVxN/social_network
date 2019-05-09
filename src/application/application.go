package application

import (
	"database/sql"
	log "social_network/src/logger"
)

var (
	Database *sql.DB

	ComLog *log.Logger
	DBlog  *log.Logger
)

func init() {
	ComLog = log.Init("common.log")
	DBlog = log.Init("database.log")
}
