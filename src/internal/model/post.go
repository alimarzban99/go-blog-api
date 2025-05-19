package model

type Post struct {
	BaseModel
	UserID      uint
	CategoryID  uint
	Title       string    `gorm:"type:varchar(250)"`
	Slug        string    `gorm:"type:varchar(250)"`
	Description *string   `gorm:"type:varchar(100)"`
	Image       string    `gorm:"type:varchar(255)"`
	Hits        uint      `gorm:"type:varchar(20);unique;not null"`
	User        *User     `gorm:"foreignKey:UserID;OnDelete:CASCADE"`
	Category    *Category `gorm:"foreignKey:CategoryID;OnDelete:CASCADE"`
}
