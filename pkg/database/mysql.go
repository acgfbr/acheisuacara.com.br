package database

import (
	"fmt"

	"acheisuacara.com.br/pkg/config"
	"acheisuacara.com.br/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(config *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.URL{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
