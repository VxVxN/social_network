package context

import (
	"database/sql"

	"github.com/VxVxN/social_network/app/log"
)

type Context struct {
	Database *sql.DB
	Log      *log.Logger
}
