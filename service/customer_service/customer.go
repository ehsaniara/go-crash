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
	ModifiedBy string
}

func (c *Customer) GetCustomer() (*models.Customer, error) {
	var customer *models.Customer

	key := fmt.Sprintf("CUSTOMER_%s", strconv.Itoa(c.ID))

	//check the cache (redis)
	data, err := redis.Get(key)
	if err != nil {
		log.Print(err)
	} else {
		err := json.Unmarshal(data, &customer)
		if err != nil {
			return nil, err
		}
		if customer.ID > 0 {
			fmt.Printf("Customer found in redis, id:%d\n", customer.ID)
			return customer, nil
		}
	}

	//if not exist
	customer, err = models.GetCustomerById(c.ID)
	if err != nil {
		return nil, err
	} else {
		fmt.Printf("Customer found in PG, id:%d\n", customer.ID)
	}

	if customer.ID == 0 {
		fmt.Printf("Customer not found in eather Redis or PG, id:%d\n", customer.ID)
		return nil, nil
	}

	err = redis.Set(key, customer, config.AppConfig.App.ObjectCashTtl)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) AddCustomer() (*models.Customer, error) {

	customer, err := models.AddNewCustomer(models.Customer{
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
	key := fmt.Sprintf("CUSTOMER_%s", strconv.Itoa(customer.ID))

	err = redis.Set(key, customer, config.AppConfig.App.ObjectCashTtl)
	if err != nil {
		return nil, err
	}

	return customer, err
}
