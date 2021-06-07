package v1

import (
	"github.com/ehsaniara/go-crash/service/customer_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Customer struct {
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Title     string `form:"title" json:"title" binding:"required"`
}

func GetCustomerById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	customerService := customer_service.Customer{ID: id}

	customer, err := customerService.GetCustomer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if customer == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, customer)
}
