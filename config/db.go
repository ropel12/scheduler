package config

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection(c *Config) (*gorm.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("unable to access database sql: %v", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to establish a good connection to the database: %v", err)
	}

	return db, nil
}
