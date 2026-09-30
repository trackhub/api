package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/trackhub/api/gorm/repository"
)

type Resolver struct {
	DB *gorm.DB
}

func (r Resolver) TrackRepository() *repository.TrackRepository {
	return repository.NewTrackRepository(r.DB)
}

func InitDB(user string, pass string, database string, host string) *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, database),
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)

	if err != nil {
		panic("unable to connect to db")
	}

	return db
}
