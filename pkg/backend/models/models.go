package models

import (
	/*blank import to make drivers available*/
	_ "database/sql"
	"fmt"
	"github.com/rfornea/library/pkg/backend/utils"
	/*blank import to make drivers available*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	/*DB is our connection to the database*/
	DB *gorm.DB
)

var DuplicateEntryStr = "Duplicate entry"

/*Connect to a database*/
func Connect(dbURL string) {
	if DB != nil {
		DB.Close()
	}
	var err error
	fmt.Println("Attempting connection to: " + dbURL)

	DB, err = gorm.Open("mysql", dbURL)
	utils.PanicOnError(err)

	DB.AutoMigrate(&Book{})
}

/*Close a database connection*/
func Close() {
	DB.Close()
}
