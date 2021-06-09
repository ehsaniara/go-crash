package customer_service

import (
	"encoding/json"
	"fmt"
	"github.com/ehsaniara/go-crash/config"
	"github.com/ehsaniara/go-crash/models"
	"github.com/ehsaniara/go-crash/pkg/redis"
	"log"
	"strconv"
)

type Customer struct {
	ID        int
	FirstName string
	LastName  string
	Title     string

	CreatedBy  string
	ModifiedBy string `json:"modifiedBy,omitempty"`
	CreatedOn  int
	ModifiedOn int
}

func (c *Customer) GetCustomer() (*Customer, error) {
	var customerModel *models.Customer

	key := fmt.Sprintf("CUSTOMER_%s", strconv.Itoa(c.ID))

	//check the cache (redis)
	data, err := redis.Get(key)
	if err != nil {
		log.Print(err)
	} else {
		err := json.Unmarshal(data, &customerModel)
		if err != nil {
			return nil, err
		}
		if customerModel.ID > 0 {
			fmt.Printf("Customer found in redis, id:%d\n", customerModel.ID)
			return customerModelToCustomer(*customerModel)
		}
	}

	//if not exist
	customerModel, err = models.GetCustomerById(c.ID)
	if err != nil {
		return nil, err
	} else {
		fmt.Printf("Customer found in PG, id:%d\n", customerModel.ID)
	}

	if customerModel.ID == 0 {
		fmt.Printf("Customer not found in eather Redis or PG, id:%d\n", customerModel.ID)
		return nil, nil
	}

	err = redis.Set(key, customerModel, config.AppConfig.App.ObjectCashTtl)
	if err != nil {
		return nil, err
	}

	return customerModelToCustomer(*customerModel)
}

func (c *Customer) AddCustomer() (*Customer, error) {

	customerModel, err := models.AddNewCustomer(models.Customer{
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		Title:      c.Title,
		CreatedBy:  c.CreatedBy,
		ModifiedBy: c.ModifiedBy,
	})
	if err != nil {
		return nil, err
	}

	//store it in the Redis
	key := fmt.Sprintf("CUSTOMER_%s", strconv.Itoa(customerModel.ID))

	err = redis.Set(key, customerModel, config.AppConfig.App.ObjectCashTtl)
	if err != nil {
		return nil, err
	}

	return customerModelToCustomer(*customerModel)
}

func customerModelToCustomer(customerModel models.Customer) (customer *Customer, err error) {

	return &Customer{
		ID:         customerModel.ID,
		FirstName:  customerModel.FirstName,
		LastName:   customerModel.LastName,
		Title:      customerModel.Title,
		CreatedBy:  customerModel.CreatedBy,
		ModifiedBy: customerModel.ModifiedBy,
		CreatedOn:  int(customerModel.CreatedOn.Unix()),
		ModifiedOn: int(customerModel.ModifiedOn.Unix()),
	}, nil
}
