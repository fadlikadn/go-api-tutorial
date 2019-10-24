package seed

import (
	"github.com/fadlikadn/go-api-tutorial/api/models"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var users = []models.User {
	models.User{
		Name:     "Steven Victor",
		Email:    "steven@gmail.com",
		Password: "password",
		Phone:    "085729789815",
		Company:  "Company 1",
		IsActive: true,
		Notes:    "",
	},
	models.User{
		Name:     "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
		Phone:    "085729789816",
		Company:  "Company 2",
		IsActive: false,
		Notes:    "",
	},
	models.User{
		Name:     "Fadlika",
		Email:    "fadlikadn@gmail.com",
		Password: "password",
		Phone:    "085729789817",
		Company:  "Company 3",
		IsActive: true,
		Notes:    "",
	},
}

var posts = []models.Post {
	models.Post{
		Title: 		"Title 1",
		Content:	"Hello World 1",
	},
	models.Post{
		Title: 		"Title 2",
		Content:	"Hello World 2",
	},
	models.Post{
		Title: 		"Title 3",
		Content:	"Hello World 3",
	},
}

var serviceTransactions = []models.ServiceTransaction{
	models.ServiceTransaction{
		ServiceDate: time.Date(2019, 8, 31, 0,0, 0, 0, time.UTC),
		InvoiceNo:   "1000",
		CustomerID:  1,
		ItemName:    "Lenovo B-470",
		DamageType:  "Mati",
		Equipment:   "Charger, Kabel Power, Kabel Data",
		Description: "Garansi 1 bulan",
		Technician:  "Agus Widodo",
		RepairType:  "Service Mainboard",
		SparePart:   "Mainboard ganti",
		Price:       650000,
		TotalPrice:  650000,
		TakenDate:   time.Date(2019, 9, 20, 0, 0, 0, 0, time.UTC),
		Status:      "Masuk",
	},
}

var customers = []models.Customer{
	models.Customer{
		Name:      "Mualifin",
		Email:     "mualifin@gmail.com",
		Phone:     "085729801987",
		Address:   "Manggisan Asri",
		Notes:     "",
	},
	models.Customer{
		Name:      "Nurafni Retno Kurniasih",
		Email:     "nurafni@gmail.com",
		Phone:     "081223564781",
		Address:   "Kalikajar Wonosobo",
		Notes:     "",
	},
	models.Customer{
		Name:      "Sigit Sambada",
		Email:     "sigit@gmail.com",
		Phone:     "081225678091",
		Address:   "Sidojoyo Wonosobo",
		Notes:     "",
	},
}

func MigrateOnly(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Customer{}, &models.ServiceTransaction{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	// Service Transaction Foreign Key
	err = db.Debug().Model(&models.ServiceTransaction{}).AddForeignKey("customer_id", "customers{id}", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.ServiceTransaction{}, &models.User{}, &models.Session{}, &models.Customer{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Session{}, &models.Customer{}, &models.ServiceTransaction{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// Post Foreign Key
	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	// Service Transaction Foreign Key
	err = db.Debug().Model(&models.ServiceTransaction{}).AddForeignKey("customer_id", "customers(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}

	for j, _ := range customers {
		err = db.Debug().Model(&models.Customer{}).Create(&customers[j]).Error
		if err != nil {
			log.Fatalf("cannot seed customers table: %v", err)
		}
	}
	for i, _ := range serviceTransactions {
		err = db.Debug().Model(&models.ServiceTransaction{}).Create(&serviceTransactions[i]).Error
		if err != nil {
			log.Fatalf("cannot seed service transactions table: %v", err)
		}
	}
}
