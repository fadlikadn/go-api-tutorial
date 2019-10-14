package models

import (
	"time"
)

type Session struct {
	Token	string		`gorm:"primary_key;type:char(43)" json:"token"`
	Data	string 		`gorm:"type:blob;not null" json:"data"`
	Expiry	time.Time	`gorm:"index:sessions_expiry_idx" json:"expiry"`
}
