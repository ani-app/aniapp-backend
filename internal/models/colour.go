package models

type Colour struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name" gorm:"not null"`
	Code string `json:"code" gorm:"not null;type:varchar(6)"`
}
