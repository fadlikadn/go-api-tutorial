package models

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/metakeule/fmtdate"
	"html"
	"strconv"
	"strings"
	"time"
)

type ServiceTransaction struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	ServiceDate	time.Time	`gorm:"not null" json:"service_date"` // Tanggal Service
	InvoiceNo	string		`gorm:"size:30; not_null" json:"invoice_no"` // Nomor Invoice
	Customer	Customer	`json:"customer"` // Pelanggan
	CustomerID	uint32		`gorm:"not_null" json:"customer_id"`
	ItemName	string		`gorm:"size:100;not null" json:"item_name"` // Nama Barang
	DamageType	string		`gorm:"size:100;not null" json:"damage_type"` // Jenis Kerusakan TODO normalize to table damage_type
	Equipment	string		`gorm:"size:100;" json:"equipment"` // Kelengkapan TODO normalize to add many-to-many relation with equipment table
	Description	string		`gorm:"size:100;" json:"description"` // Keterangan/garansi notes
	Technician	string 		`gorm:"size:40;" json:"technician"` // Teknisi TODO normalize to table technician
	RepairType	string		`gorm:"size:100;" json:"repair_type"` // Jenis Perbaikan TODO normalize to table repair_type
	SparePart	string		`gorm:"size:100;" json:"spare_part"`// Spare Part TODO normalize to add many-to-many relation with table spare_part
	Price		uint64		`json:"price"` // Harga
	TotalPrice	uint64		`json:"total_price"` // Total Harga
	TakenDate	time.Time	`json:"taken_date"` // Tanggal Pengambilan
	Status		string		`json:"status"` // Status Perbaikan TODO normalize to table service_status
}

func (s *ServiceTransaction) Prepare() {
	s.ID = 0
	s.InvoiceNo = html.EscapeString(strings.TrimSpace(s.InvoiceNo))
	s.Customer = Customer{}
	s.ItemName = html.EscapeString(strings.TrimSpace(s.ItemName))
	s.DamageType = html.EscapeString(strings.TrimSpace(s.DamageType))
	s.Equipment = html.EscapeString(strings.TrimSpace(s.Equipment))
	s.Description = html.EscapeString(strings.TrimSpace(s.Description))
	s.Technician = html.EscapeString(strings.TrimSpace(s.Technician))
	s.RepairType = html.EscapeString(strings.TrimSpace(s.RepairType))
	s.SparePart = html.EscapeString(strings.TrimSpace(s.SparePart))
	s.Status = html.EscapeString(strings.TrimSpace(s.Status))
}

func (s *ServiceTransaction) Validate() error {
	if s.InvoiceNo == "" {
		return errors.New("Required Invoice No")
	}
	if s.ItemName == "" {
		return errors.New("Required Item Name")
	}
	if s.DamageType == "" {
		return errors.New("Required Damage Type")
	}
	return nil
}

func (s *ServiceTransaction) SaveServiceTransaction(db *gorm.DB) (*ServiceTransaction, error) {
	var err error
	err = db.Debug().Model(&ServiceTransaction{}).Create(&s).Error
	if err != nil {
		return &ServiceTransaction{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&Customer{}).Where("id = ?", s.CustomerID).Take(&s.Customer).Error
		if err != nil {
			return &ServiceTransaction{}, err
		}
	}
	return s, nil
}

func (s *ServiceTransaction) FindAllServiceTransactions(db *gorm.DB) (*[]ServiceTransaction, error) {
	var err error
	serviceTransactions := []ServiceTransaction{}
	err = db.Debug().Model(&ServiceTransaction{}).Limit(100).Find(&serviceTransactions).Error
	if err != nil {
		return &[]ServiceTransaction{}, err
	}
	if len(serviceTransactions) > 0 {
		for i, _ := range serviceTransactions {
			err := db.Debug().Model(&Customer{}).Where("id = ?", serviceTransactions[i].CustomerID).Take(&serviceTransactions[i].Customer).Error
			if err != nil {
				return &[]ServiceTransaction{}, err
			}
		}
	}
	return &serviceTransactions, nil
}

func (s *ServiceTransaction) FindServiceTransactionByID(db *gorm.DB, pid uint32) (*ServiceTransaction, error) {
	var err error
	err = db.Debug().Model(&ServiceTransaction{}).Where("id = ?", pid).Take(&s).Error
	if err != nil {
		return &ServiceTransaction{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&Customer{}).Where("id = ?", s.CustomerID).Take(&s.Customer).Error
		if err != nil {
			return &ServiceTransaction{}, err
		}
	}
	return s, nil
}

func (s *ServiceTransaction) UpdateServiceTransaction(db *gorm.DB, pid uint32) (*ServiceTransaction, error) {
	var err error
	db = db.Debug().Model(&ServiceTransaction{}).Where("id = ?", pid).Take(&ServiceTransaction{}).UpdateColumns(
			map[string]interface{} {
				"service_date": s.ServiceDate,
				"invoice_no": s.InvoiceNo,
				"customer_id": s.CustomerID,
				"item_name": s.ItemName,
				"damage_type": s.DamageType,
				"equipment": s.Equipment,
				"description": s.Description,
				"technician": s.Technician,
				"repair_type": s.RepairType,
				"spare_part": s.SparePart,
				"price": s.Price,
				"total_price": s.TotalPrice,
				"taken_date": s.TakenDate,
				"status": s.Status,
			},
	)
	err = db.Debug().Model(&ServiceTransaction{}).Where("id = ?", pid).Take(&s).Error
	if err != nil {
		return &ServiceTransaction{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&ServiceTransaction{}).Where("id = ?", s.CustomerID).Take(&s.Customer).Error
		if err != nil {
			return &ServiceTransaction{}, err
		}
	}

	return s, nil
}

func (s *ServiceTransaction) DeleteServiceTransaction(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&ServiceTransaction{}).Where("id = ?", uid).Take(&ServiceTransaction{}).Delete(&ServiceTransaction{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *ServiceTransaction) UnmarshalJSON(j []byte, customer int) error {
	var rawStrings map[string]string
	
	err := json.Unmarshal(j, &rawStrings)
	if err != nil {
		return err
	}

	if customer != 0 {
		s.CustomerID = uint32(customer)
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "service_date" {
			sd, err := fmtdate.Parse("DD/MM/YYYY", v)
			if err != nil {
				return err
			}
			s.ServiceDate = sd
		}

		if strings.ToLower(k) == "invoice_no" {
			s.InvoiceNo = v
		}

		if strings.ToLower(k) == "item_name" {
			s.ItemName = v
		}

		if strings.ToLower(k) == "damage_type" {
			s.DamageType = v
		}

		if strings.ToLower(k) == "equipment" {
			s.Equipment = v
		}

		if strings.ToLower(k) == "description" {
			s.Description = v
		}

		if strings.ToLower(k) == "technician" {
			s.Technician = v
		}

		if strings.ToLower(k) == "repair_type" {
			s.RepairType = v
		}

		if strings.ToLower(k) == "spare_part" {
			s.SparePart = v
		}

		if strings.ToLower(k) == "price" {
			s.Price, _ = strconv.ParseUint(v, 10, 64)
		}

		if strings.ToLower(k) == "total_price" {
			s.TotalPrice, _ = strconv.ParseUint(v, 10, 64)
		}

		if strings.ToLower(k) == "invoice_no" {
			s.InvoiceNo = v
		}

		if strings.ToLower(k) == "taken_date" {
			td, err := fmtdate.Parse("DD/MM/YYYY", v)
			if err != nil {
				return err
			}
			s.TakenDate = td
		}

		if strings.ToLower(k) == "status" {
			s.Status = v
		}
	}

	return nil
}