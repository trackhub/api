package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import(
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Resolver struct{
	DB *gorm.DB
}

func InitDB(user string, pass string, database string, host string) *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, database),
		),
		&gorm.Config{},
	)

	if err != nil {
		panic("unable to connect to db")
	}

	return db
}
