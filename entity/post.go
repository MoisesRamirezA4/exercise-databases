package entity

type Post struct {
	ID   int `gorm:"primaryKey" json:"id"`
	Name string
}
