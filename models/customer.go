package models

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	Model
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	Title      string `json:"title" validate:"required"`
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

	customer.CreatedOn = int(time.Now().Unix())

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
