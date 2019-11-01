package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"html"
	"strings"
	"time"
)

type Customer struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	UUID		string		`gorm:"unique;"`
	Name		string		`gorm:"size:255;not null;unique" json:"name"`
	Email		string		`gorm:"size:255;unique" json:"email"`
	Phone		string		`gorm:"size:100;" json:"phone"`
	Address		string		`gorm:"size:255;" json:"address"`
	Notes		string		`gorm:"size:255" json:"notes"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Customer) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4().String()
	return scope.SetColumn("UUID", id)
}

func (c *Customer) Prepare() {
	c.ID = 0
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.Email = html.EscapeString(strings.TrimSpace(c.Email))
	c.Phone = html.EscapeString(strings.TrimSpace(c.Phone))
	c.Address = html.EscapeString(strings.TrimSpace(c.Address))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Customer) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if c.Name == "" {
			return errors.New("Required Name")
		}
		//if c.Phone == "" {
		//	return errors.New("Required Phone")
		//}
		//if c.Address == "" {
		//	return errors.New("Required Address")
		//}
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	default:
		if c.Name == "" {
			return errors.New("Required Name")
		}
		//if c.Phone == "" {
		//	return errors.New("Required Phone")
		//}
		//if c.Address == "" {
		//	return errors.New("Required Address")
		//}
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	}
}

func (c *Customer) SaveCustomer(db *gorm.DB) (*Customer, error) {
	var err error
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}

func (c *Customer) FindAllCustomer(db *gorm.DB) (*[]Customer, error) {
	var err error
	customers := []Customer{}
	err = db.Debug().Model(&Customer{}).Find(&customers).Error
	if err != nil {
		return &[]Customer{}, err
	}
	return &customers, err
}

func (c *Customer) FindCustomerByID(db *gorm.DB, uid uint32) (*Customer, error) {
	var err error
	err = db.Debug().Model(Customer{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Customer{}, errors.New("Customer Not Found")
	}
	return c, err
}

func (c *Customer) FindCustomerByName(db *gorm.DB, name string) (*[]Customer, error) {
	var err error
	customers := []Customer{}
	err = db.Debug().Model(&Customer{}).Where("name LIKE ?", "%" + name + "%").Find(&customers).Error
	if err != nil {
		return &[]Customer{}, err
	}
	return &customers, err
}

func (c *Customer) UpdateCustomer(db *gorm.DB, uid uint32) (*Customer, error) {
	db = db.Debug().Model(&Customer{}).Where("id = ?", uid).Take(&Customer{}).UpdateColumns(
		map[string]interface{}{
			"name": c.Name,
			"email": c.Email,
			"phone": c.Phone,
			"address": c.Address,
			"notes": c.Notes,
			"updated_at": time.Now(),
		},
	)

	if db.Error != nil {
		return &Customer{}, db.Error
	}
	// This is display the updated Customer
	err := db.Debug().Model(&Customer{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}

func (c *Customer) DeleteCustomer(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&Customer{}).Where("id = ?", uid).Take(&Customer{}).Delete(&Customer{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

