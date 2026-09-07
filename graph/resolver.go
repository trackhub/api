package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Resolver struct{
	DB *gorm.DB
}

func InitDB() *gorm.DB {
	// @TODO move to env
	dsn := "gps:1@tcp(sql:3306)/gps?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("unable to connect to db")
	}

	return db
}
