package database

import (
	"fmt"
	"github.com/alimarzban99/go-blog-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {

	once.Do(func() {
		user := config.Config.Database.User
		pass := config.Config.Database.Password
		host := config.Config.Database.Host
		port := config.Config.Database.Port
		name := config.Config.Database.Name

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, pass, host, port, name)

		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

	})
	return db
}
