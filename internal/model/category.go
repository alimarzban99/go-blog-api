package model

type Category struct {
	BaseModel
	Title string `gorm:"type:varchar(250)"`
	Posts *[]Post
}
