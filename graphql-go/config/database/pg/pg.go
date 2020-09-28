package pg

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// New create new postgres database connection
// it return postgres connection and cleanup postgres connection
func New(host, port, user, password, dbName string) (*gorm.DB, func()) {
	connectionString := fmt.Sprintf("host = %s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("gorm:table_options", "collation_connection=utf8_general_ci")

	return db, func() {
		err := db.Close()
		if err != nil {
			log.Println("Failed to close DB by error", err)
		}
	}
}
