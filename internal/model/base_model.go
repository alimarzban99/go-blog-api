package model

import (
	"github.com/alimarzban99/go-blog-api/internal/enums"
	"github.com/alimarzban99/go-blog-api/pkg/database"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Status    enums.Status   `gorm:"type:enum('active','inactive','banned');default:'active'"`
}

func Starter() {
	db := database.GetDB()
	err := db.AutoMigrate(&User{}, &VerificationCode{}, &Token{}, &Post{}, &Category{})
	if err != nil {
		return
	}
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
