package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Dialector for open connection based on driver
// Uncomment the code based on your database driver
// Run go get gorm.io/driver/{driver_name} to install the driver
// Example: go get gorm.io/driver/mysql and import "gorm.io/driver/mysql"
var dialectors = func(driver, dsn string) gorm.Dialector {
	return map[string]gorm.Dialector{
		"mysql": mysql.Open(dsn),
	}[driver]
}

// Connect and return *gorm.DB connection
func ConnectGorm(driver string, dbURL string) *gorm.DB {
	dialector := dialectors(driver, dbURL)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error", err)
	}

	fmt.Println("Database connected successfully")
	return db
}
