package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/leekchan/accounting"
	"github.com/metakeule/fmtdate"
	uuid "github.com/satori/go.uuid"
	"github.com/unidoc/unipdf/v3/creator"
	"html"
	"os"
	"strconv"
	"strings"
	"time"
)

type ServiceTransaction struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	UUID		string		`gorm:"unique;"`
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
	AdditionalItems	[]AdditionalItem	`gorm:"foreignkey:STId" json:"additional_items"`
}

type ServiceTransactionStatus struct {
	New 		int		`json:"new"`
	InProgress	int		`json:"in_progress"`
	Completed	int		`json:"completed"`
}

func (s *ServiceTransaction) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4().String()
	return scope.SetColumn("UUID", id)
}

func (s *ServiceTransaction) Prepare() {
	//s.ID = 0
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

func (s *ServiceTransaction) NumberByStatus(db *gorm.DB, status string) (int, error) {
	var err error
	var count int
	err = db.Debug().Model(&ServiceTransaction{}).Where("status = ?", status).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
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

			// Additional Item list
			err = db.Debug().Model(&AdditionalItem{}).Where("st_id = ?", serviceTransactions[i].ID).Find(&serviceTransactions[i].AdditionalItems).Error
			if err != nil {
				fmt.Println("error during get additional items query")
				fmt.Println(err)
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

	// Additional Item list
	err = db.Debug().Model(&AdditionalItem{}).Where("st_id = ?", s.ID).Find(&s.AdditionalItems).Error
	if err != nil {
		fmt.Println("error during get additional items query")
		fmt.Println(err)
		return &ServiceTransaction{}, err
	}

	return s, nil
}

func (s *ServiceTransaction) FindServiceTransactionByInvoiceNo(db *gorm.DB, invoiceNo string) (*ServiceTransaction, error) {
	var err error
	err = db.Debug().Model(&ServiceTransaction{}).Where("invoice_no = ?", invoiceNo).Take(&s).Error
	if err != nil {
		return &ServiceTransaction{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&Customer{}).Where("id = ?", s.CustomerID).Take(&s.Customer).Error
		if err != nil {
			return &ServiceTransaction{}, err
		}
	}

	// Additional Item list
	err = db.Debug().Model(&AdditionalItem{}).Where("st_id = ?", s.ID).Find(&s.AdditionalItems).Error
	if err != nil {
		fmt.Println("error during get additional items query")
		fmt.Println(err)
		return &ServiceTransaction{}, err
	}

	return s, nil
}

func (s *ServiceTransaction) FindServiceTransactionByUUID(db *gorm.DB, uuid string) (*ServiceTransaction, error) {
	var err error
	err = db.Debug().Model(&ServiceTransaction{}).Where("uuid = ?", uuid).Take(&s).Error
	if err != nil {
		return &ServiceTransaction{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&Customer{}).Where("id = ?", s.CustomerID).Take(&s.Customer).Error
		if err != nil {
			return &ServiceTransaction{}, err
		}
	}

	// Additional Item list
	err = db.Debug().Model(&AdditionalItem{}).Where("st_id = ?", s.ID).Find(&s.AdditionalItems).Error
	if err != nil {
		fmt.Println("error during get additional items query")
		fmt.Println(err)
		return &ServiceTransaction{}, err
	}

	return s, nil
}

func (s *ServiceTransaction) UpdateServiceTransaction(db *gorm.DB, pid uint32) (*ServiceTransaction, error) {
	var err error
	fmt.Println(pid)
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
		fmt.Println("error when get service transaction update")
		fmt.Println(err)
		return &ServiceTransaction{}, err
	}
	//if s.ID != 0 {
	//	err = db.Debug().Model(&Customer{}).Where("id = ?", s.CustomerID).Take(&s.Customer).Error
	//	if err != nil {
	//		fmt.Println("error when get customer")
	//		return &ServiceTransaction{}, err
	//	}
	//}

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

// TODO Implement Create Invoice for Service Transaction
func (s *ServiceTransaction) CreateInvoice(uuid string, db *gorm.DB, c *creator.Creator, logoPath string) (*creator.Invoice, error) {
	// Create a new invoice
	invoice := c.NewInvoice()
	ac := accounting.Accounting{Symbol: "Rp ", Precision:0, Thousand: ".", Decimal: ","}

	if (uuid == "") {
		// Create an instance of Logo used as a header for the invoice
		// If the image is not stored locally, you can use NewImageFromData to generate it byte array
		/*logo, err := c.NewImageFromData([]byte("Logo Company"))
		if err != nil {
			return nil, err
		}*/

		// Set invoice logo
		//invoice.SetLogo(logo)

		// Set invoice information
		invoice.SetNumber("0001")
		invoice.SetDate("01/11/2019")
		invoice.SetDueDate("10/11/2019")
		invoice.AddInfo("Payment terms", "Due on receipt")
		invoice.AddInfo("Paid", "No")

		// Set invoice address
		invoice.SetSellerAddress(&creator.InvoiceAddress{
			Name:    "John Doe",
			Street:  "8 Elm Street",
			Zip:     "Cambridge",
			City:    "56351",
			Country: "Indonesia",
			Phone:   "081-xxx-xxx-987",
			Email:   "johndoe@gmail.com",
		})

		invoice.SetBuyerAddress(&creator.InvoiceAddress{
			Name:    "Jane Doe",
			Street:  "9 Elm Street",
			Zip:     "56372",
			City:    "London",
			Country: "Indonesia",
			Phone:   "081804xxx897",
			Email:   "janedoe@gmail.com",
		})

		// Add products to invoice
		for i := 1; i < 6; i++ {
			invoice.AddLine(
				fmt.Sprintf("Test product #%d", 1),
				"1",
				strconv.Itoa((i-1) * 7),
				strconv.Itoa((i + 4) * 3),
			)
		}

		// Set invoice totals
		invoice.SetSubtotal("Rp 1.000.000")
		invoice.AddTotalLine("Tax (10%)", "Rp 100.000")
		invoice.AddTotalLine("Shipping", "Rp 50.000")
		invoice.SetTotal("Rp 1.150.000")

		// Set invoice content sections
		invoice.SetNotes("Notes", "Thank you for your business")
		invoice.SetTerms("Terms and conditions", "Full refund for 60 days after purchase")

		return invoice, nil
	} else {
		// Create an instance of Logo used as a header for the invoice
		// If the image is not stored locally, you can use NewImageFromData to generate it byte array
		//logo, err := c.NewImageFromData([]byte("assets/image/company_logo.png"))
		logo, err := c.NewImageFromFile("assets/image/company_logo.png")
		logo.ScaleToWidth(500)
		if err != nil {
			return nil, err
		}

		// Set invoice logo
		invoice.SetLogo(logo)

		serviceTransaction := ServiceTransaction{}
		serviceTransactionFound, err := serviceTransaction.FindServiceTransactionByUUID(db, uuid)
		if err != nil {
			return nil, err
		}

		invoice.SetNumber(serviceTransactionFound.InvoiceNo)
		invoice.SetDate(fmtdate.Format("DD/MM/YYYY", serviceTransactionFound.ServiceDate))
		invoice.SetDueDate(fmtdate.Format("DD/MM/YYYY", serviceTransactionFound.TakenDate))
		invoice.AddInfo("Tipe Reparasi", serviceTransactionFound.RepairType)
		invoice.AddInfo("Kelengkapan", serviceTransactionFound.Equipment)
		invoice.AddInfo("Kerusakan", serviceTransactionFound.DamageType)
		invoice.AddInfo("Dibayar", "Belum")

		// Set invoice address
		invoice.SetSellerAddress(&creator.InvoiceAddress{
			Name:    os.Getenv("COMPANY_NAME"),
			Street:  os.Getenv("COMPANY_ADDRESS"),
			//Zip:     "Cambridge",
			City:    os.Getenv("COMPANY_CITY"),
			//Country: "Indonesia",
			Phone:   os.Getenv("COMPANY_PHONE"),
			//Email:   "johndoe@gmail.com",
		})

		// Set buyer name
		invoice.SetBuyerAddress(&creator.InvoiceAddress{
			Name:    serviceTransactionFound.Customer.Name,
			Street:  serviceTransaction.Customer.Address,
			//Zip:     "-",
			//City:    "-",
			Country: "Indonesia",
			Phone:   serviceTransaction.Customer.Phone,
			Email:   serviceTransaction.Customer.Email,
		})

		// Add transaction to invoice
		invoice.AddLine(
			serviceTransactionFound.ItemName,
			"1",
			ac.FormatMoney(int(serviceTransactionFound.Price)),
			ac.FormatMoney(int(serviceTransactionFound.Price)),
		)

		for i := 0; i < len(serviceTransactionFound.AdditionalItems); i++ {
			invoice.AddLine(
				serviceTransactionFound.AdditionalItems[i].Name,
				"1",
				ac.FormatMoney(int(serviceTransactionFound.AdditionalItems[i].Cost)),
				ac.FormatMoney(int(serviceTransactionFound.AdditionalItems[i].Cost)),
			)
		}

		// Set invoice totals
		invoice.SetSubtotal(ac.FormatMoney(int(serviceTransactionFound.TotalPrice)))
		invoice.SetTotal(ac.FormatMoney(int(serviceTransactionFound.TotalPrice)))

		// Set invoice content sections
		invoice.SetNotes("Notes", "Thank you for your business")
		invoice.SetTerms("Terms and conditions", "Garansi 1 bulan")

		return invoice, nil
	}

}