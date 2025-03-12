package database

import (
	"database/sql"

	"gorm.io/gorm"
)

var (
	SqlDb *sql.DB
	Db    *gorm.DB
)
