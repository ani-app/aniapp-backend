package models

import (
	"gorm.io/gorm"
	"time"
)

type Pet struct {
	gorm.Model
	Name       string   `gorm:"not null"`
	CustomerID uint     `gorm:"not null"`
	Customer   Customer `gorm:"embedded"`
	Birthday   time.Time
	PhotoUrl   string
	IdNumber   string
	BreedID    uint   `gorm:"not null"`
	Breed      Breed  `gorm:"embedded"`
	ColourID   uint   `gorm:"not null"`
	Colour     Colour `gorm:"embedded"`
	GenderID   uint   `gorm:"not null"`
	Gender     Gender `gorm:"embedded"`
}

type PetDTO struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	CustomerID uint      `json:"customer_id,omitempty"`
	Customer   Customer  `json:"customer,omitempty"`
	Birthday   time.Time `json:"birthday"`
	PhotoUrl   string    `json:"photo_url"`
	IdNumber   string    `json:"id_number"`
	BreedID    uint      `json:"breed_id,omitempty"`
	Breed      Breed     `json:"breed,omitempty"`
	ColourID   uint      `json:"colour_id,omitempty"`
	Colour     Colour    `json:"colour,omitempty"`
	GenderID   uint      `json:"gender_id,omitempty"`
	Gender     Gender    `json:"gender,omitempty"`
}

func (p Pet) toPetDTO() PetDTO {
	return PetDTO{
		ID:         p.ID,
		Name:       p.Name,
		CustomerID: p.CustomerID,
		Customer:   p.Customer,
		Birthday:   p.Birthday,
		PhotoUrl:   p.PhotoUrl,
		IdNumber:   p.IdNumber,
		Breed:      p.Breed,
		Colour:     p.Colour,
		Gender:     p.Gender,
	}
}

func (pd PetDTO) toPet() Pet {
	return Pet{
		Model:      gorm.Model{ID: pd.ID},
		Name:       pd.Name,
		CustomerID: pd.CustomerID,
		Customer:   pd.Customer,
		Birthday:   pd.Birthday,
		PhotoUrl:   pd.PhotoUrl,
		IdNumber:   pd.IdNumber,
		BreedID:    pd.BreedID,
		Breed:      pd.Breed,
		ColourID:   pd.ColourID,
		Colour:     pd.Colour,
		GenderID:   pd.GenderID,
		Gender:     pd.Gender,
	}
}

func ToPetDTOs(pets []Pet) []PetDTO {
	petDTOs := make([]PetDTO, len(pets))
	for i, p := range pets {
		petDTOs[i] = p.toPetDTO()
	}
	return petDTOs
}
