package storage

import (
	"database/sql"
	"github.com/nats-io/nats.go"
)

type Storage struct {
	db     *sql.DB
	broker nats.Conn
}
