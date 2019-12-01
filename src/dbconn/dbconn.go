package dbconn

import (
	"github.com/jinzhu/gorm"

	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//OpenConnection open a connection to databse
func OpenConnection() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test2 password=amin")
	if err != nil {
		return nil, err
	}
	return db, err
}
