package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type Repair struct {
	ID uint64	`gorm:"primary_key;auto_increment" json:"id"`
	Client Customer	`json:"client"`
	ClientID	uint32	`gorm:"not_null" json:"client_id"`
	Nota		string		`gorm:"size:255" json:"nota"`
	Title		string		`gorm:"size:255;not_null;" json:"title"`
	Description	string		`gorm:"size:255" json:"description"`
	RepairDate	time.Time	`json:"repair_date"`
	Status		string		`form:"size:50" json:"status"`
	CreatedAt 	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (r *Repair) Prepare() {
	r.ID = 0
	r.Nota = html.EscapeString(strings.TrimSpace(r.Nota))
	r.Title = html.EscapeString(strings.TrimSpace(r.Title))
	r.Description = html.EscapeString(strings.TrimSpace(r.Description))
	r.Status = html.EscapeString(strings.TrimSpace(r.Status))
}

func (r *Repair) Validate() error {
	if r.Title == "" {
		return errors.New("Required Title")
	}
	if r.Status == "" {
		return errors.New("Required Title")
	}
	return nil
}



