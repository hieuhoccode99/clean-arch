package models

type Article struct {
	BaseModel
	Title string `json:"title" gorm:"column:title"`
}
