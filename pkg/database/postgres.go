package database

import (
	"fmt"
	"github.com/alimarzban99/go-blog-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func InitDb() error {
	var err error
	user := config.Config.Database.User
	pass := config.Config.Database.Password
	host := config.Config.Database.Host
	port := config.Config.Database.Port
	name := config.Config.Database.Name
	sslMode := config.Config.Database.SSLMode
	maxIdleConns := config.Config.Database.MaxIdleConns
	maxOpenConns := config.Config.Database.MaxOpenConns
	connMaxLifetime := config.Config.Database.ConnMaxLifetime

	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		host, port, user, pass,
		name, sslMode)

	db, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := db.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(maxIdleConns)
	sqlDb.SetMaxOpenConns(maxOpenConns)
	sqlDb.SetConnMaxLifetime(connMaxLifetime)

	log.Println("Db connection established")
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDb() {
	con, _ := db.DB()
	err := con.Close()
	if err != nil {
		log.Println(err.Error())
	}
}
