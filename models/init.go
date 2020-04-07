package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitialDatabase() {
	var connectErr error

	// connectionString := fmt.Sprintf("host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	// connectionString := fmt.Sprintf("host=35.240.191.214 port=5432 user=postgres dbname=fillgoods-lab password=IkDD43cIamwavO7s")
	connectionString := fmt.Sprintf("host=127.0.0.1 port=5432 user=user1 dbname=db password=1234 sslmode = disable")
	DB, connectErr = gorm.Open("postgres", connectionString)
	if connectErr != nil {
		fmt.Println(connectErr)
		return
	}

	autoMigrateDatabase()
	// defer db.Close()
	fmt.Println("Database connection successfully...")
}

func autoMigrateDatabase() {
	DB.AutoMigrate(&User{})
}
