package adapter

import (
	"database/sql"
)

// Database abstrai a conexão com o db
type Database interface {
	Connect() (*sql.DB, error)
}
