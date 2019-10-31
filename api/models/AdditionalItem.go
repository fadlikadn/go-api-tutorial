package models

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strconv"
	"strings"
)

type AdditionalItem struct {
	ID		uint32	`gorm:"primary_key;autoincrement" json:"id"`
	Name	string	`gorm:"size:100;not null" json:"name"`	// Name
	Notes	string	`gorm:"size:150;" json:"notes"`	// Notes
	Cost	uint64	`gorm:"not null" json:"cost"`	// Cost Additional item
	STId	uint32	`gorm:"not null" json:"st_id"` //Service Transaction ID
}

func (a *AdditionalItem) Prepare() {
	a.ID = 0
	a.Name = html.EscapeString(strings.TrimSpace(a.Name))
	a.Notes = html.EscapeString(strings.TrimSpace(a.Notes))
}

func (a *AdditionalItem) Validate() error {
	if a.Name == "" {
		return errors.New("Required name")
	}
	return nil
}

func (a *AdditionalItem) SaveAdditionalItem(db *gorm.DB) (*AdditionalItem, error) {
	var err error
	err = db.Debug().Model(&AdditionalItem{}).Create(&a).Error
	if err != nil {
		return &AdditionalItem{}, err
	}
	return a, nil
}

func (a *AdditionalItem) FindAllAdditionalItemBySTId(db *gorm.DB, stid uint32) (*[]AdditionalItem, error) {
	var err error
	additionalItems := []AdditionalItem{}
	err = db.Debug().Model(&AdditionalItem{}).Where("st_id = ?", stid).Find(&additionalItems).Error
	if err != nil {
		return &[]AdditionalItem{}, err
	}

	return &additionalItems, nil
}

func (a *AdditionalItem) DeleteAllAdditionalItemBySTid(db *gorm.DB, stid uint32) (int64, error) {
	db = db.Debug().Model(&AdditionalItem{}).Where("st_id = ?", stid).Delete(&AdditionalItem{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (a *AdditionalItem) UnmarshalJSON(j []byte, serviceTransactionId int) error {
	var rawStrings map[string]string

	err := json.Unmarshal(j, &rawStrings)
	if err != nil {
		return err
	}

	if serviceTransactionId != 0 {
		a.STId = uint32(serviceTransactionId)
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "name" {
			a.Name = v
		}
		if strings.ToLower(k) == "notes" {
			a.Notes = v
		}
		if strings.ToLower(k) == "cost" {
			a.Cost, _ = strconv.ParseUint(v, 10, 64)
		}
	}

	return nil
}






