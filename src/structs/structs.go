package structs

import (
	"github.com/jinzhu/gorm"
)

//Register struct
type Register struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//Login struct
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//User struct
type User struct {
	gorm.Model
	ID        int    `gorm:"AUTO_INCREMENT"`
	FirstName string `gorm:"type:varchar(100);column:first_name"`
	LastName  string `gorm:"type:varchar(100);column:last_name"`
	Email     string `gorm:"primary_key;type:varchar(100);column:email_id"`
	Password  string `gorm:"type:varchar(100);column:password"`
}
