package v1

import (
	"github.com/ehsaniara/go-crash/pkg/log"
	"github.com/ehsaniara/go-crash/service/customer_service"
	"github.com/ehsaniara/go-crash/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

type AddCustomer struct {
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Title     string `form:"title" json:"title" binding:"required"`
}

func AddNewCustomer(c *gin.Context) {
	var (
		addCustomer AddCustomer
	)

	err := c.BindJSON(&addCustomer)
	if err != nil {
		return
	}

	claims, _ := c.Get("claims")

	var claimsVar = claims.(util.Claims)

	log.Log.Debugf("Create username: %v", claimsVar.Username)

	customerService := customer_service.Customer{
		FirstName:  addCustomer.FirstName,
		LastName:   addCustomer.LastName,
		Title:      addCustomer.Title,
		CreatedBy:  claimsVar.Username,
		ModifiedBy: "",
	}
	customer, err := customerService.AddCustomer()

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, customer)
}
