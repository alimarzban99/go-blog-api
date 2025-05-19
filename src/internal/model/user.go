package model

import (
	"database/sql"
	"github.com/alimarzban99/go-blog-api/internal/enums"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	FirstName sql.NullString `gorm:"type:varchar(100);null"`
	LastName  sql.NullString `gorm:"type:varchar(100);null"`
	Mobile    string         `gorm:"type:varchar(20);unique;not null"`
	IsAdmin   bool           `gorm:"type:boolean;not null;default:false"`
	Email     *string        `gorm:"type:varchar(100);unique"`
	Posts     *[]Post
}

func ActiveUser(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", enums.Active)
}
