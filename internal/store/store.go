package store

import (
	"Identity/cmd/config"
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewMSSQLStorage(conf config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	return db, err
}
