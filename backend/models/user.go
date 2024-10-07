package models

import "database/sql"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Fullname sql.NullString
	Password string
}
