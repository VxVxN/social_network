package context

import (
	"database/sql"

	"github.com/VxVxN/social_network/internal/log"
)

type Context struct {
	Database *sql.DB
	Log      *log.Logger
}
