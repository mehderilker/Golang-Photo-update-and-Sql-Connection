package models

import (
	"time"
	"database/sql"
)

type Photo struct {
	ID int
	Pname sql.NullString
	Date time.Time
	Datestr string
	Status bool
}
