package models

import "gorm.io/gorm"

type CustomerProfile struct {
	gorm.Model
	Firstname       string `gorm:"not null"`
	Lastname        string `gorm:"not null"`
	ProfilePhotoUrl string
}

type CustomerProfileDTO struct {
	ID              uint   `json:"id"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	ProfilePhotoUrl string `json:"profile_photo_url"`
}

func (cp CustomerProfile) toCustomerProfileDTO() CustomerProfileDTO {
	return CustomerProfileDTO{
		ID:              cp.ID,
		Firstname:       cp.Firstname,
		Lastname:        cp.Lastname,
		ProfilePhotoUrl: cp.ProfilePhotoUrl,
	}
}

func (cpd CustomerProfileDTO) toCustomerProfile() CustomerProfile {
	return CustomerProfile{
		Model:           gorm.Model{ID: cpd.ID},
		Firstname:       cpd.Firstname,
		Lastname:        cpd.Lastname,
		ProfilePhotoUrl: cpd.ProfilePhotoUrl,
	}
}

func ToCustomerProfileDTOs(customerProfiles []CustomerProfile) []CustomerProfileDTO {
	customerProfileDTOs := make([]CustomerProfileDTO, len(customerProfiles))
	for i, p := range customerProfiles {
		customerProfileDTOs[i] = p.toCustomerProfileDTO()
	}
	return customerProfileDTOs
}
