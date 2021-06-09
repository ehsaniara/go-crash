package models

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	Model
	FirstName  string `json:"firstName" validate:"required" gorm:"not null"`
	LastName   string `json:"lastName" validate:"required" gorm:"not null"`
	Title      string `json:"title" validate:"required" gorm:"not null"`
	CreatedBy  string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
}

func GetCustomerById(id int) (*Customer, error) {
	var customer Customer
	err := db.Where("id = ?", id).First(&customer).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &customer, nil
}

func AddNewCustomer(customer Customer) (*Customer, error) {

	customer.CreatedOn = time.Now()
	customer.ModifiedOn = time.Now()

	db.Save(&customer)

	//if err := db.Create(&customer).Error; err != nil {
	//	return err, GetCustomerById()
	//}

	res, err := GetCustomerById(customer.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
