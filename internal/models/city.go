package models

type City struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CountryID uint    `json:"country_id,omitempty" gorm:"not null"`
	Country   Country `json:"country,omitempty" gorm:"embedded"`
	Name      string  `json:"name" gorm:"not null"`
}
