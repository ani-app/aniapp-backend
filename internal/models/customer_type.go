package models

import "gorm.io/gorm"

type CustomerType struct {
	gorm.Model
	Name string `gorm:"not null"`
	Flag int    `gorm:"not null"`
}

type CustomerTypeDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Flag int    `json:"flag"`
}

func (ct CustomerType) toCustomerTypeDTO() CustomerTypeDTO {
	return CustomerTypeDTO{
		ID:   ct.ID,
		Name: ct.Name,
		Flag: ct.Flag,
	}
}

func (ctd CustomerTypeDTO) toCustomerType() CustomerType {
	return CustomerType{
		Model: gorm.Model{ID: ctd.ID},
		Name:  ctd.Name,
		Flag:  ctd.Flag,
	}
}

func ToCustomerTypeDTOs(customerTypes []CustomerType) []CustomerTypeDTO {
	customerTypeDTOs := make([]CustomerTypeDTO, len(customerTypes))
	for i, t := range customerTypes {
		customerTypeDTOs[i] = t.toCustomerTypeDTO()
	}
	return customerTypeDTOs
}
