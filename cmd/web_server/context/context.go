package context

import (
	"database/sql"
	"social_network/internal/log"
)

type Context struct {
	Database *sql.DB
	Log      *log.Logger
}
