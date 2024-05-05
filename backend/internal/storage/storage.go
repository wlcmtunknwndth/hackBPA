package storage

import (
	"database/sql"
	"github.com/nats-io/nats.go"
	"time"
)

type Storage struct {
	db     *sql.DB
	broker nats.Conn
}

type Event struct {
	Id           uint      `json:"id"`
	Price        uint      `json:"price"`
	Restrictions uint      `json:"restrictions"`
	Date         time.Time `json:"date"`
	Location     string    `json:"location"`
	Name         string    `json:"name"`
	ImgPath      string    `json:"img_path"`
	Description  string    `json:"description"`
	Features     Features  `json:"features"`
}

type Features struct {
	Disability bool `json:"disability"`
	Deaf       bool `json:"deaf"`
	Blind      bool `json:"blind"`
	Neural     bool `json:"neural"`
}
