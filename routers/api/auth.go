package api

import (
	"github.com/ehsaniara/go-crash/service/auth_service"
	"github.com/ehsaniara/go-crash/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {

	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authService := auth_service.Auth{Username: login.Username, Password: login.Password}
	isExist, err := authService.Check()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if !isExist {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	token, err := util.GenerateToken(login.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
