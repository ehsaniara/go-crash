package util

import (
	"github.com/ehsaniara/go-crash/config"
	"github.com/ehsaniara/go-crash/pkg/log"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		// handle error
		log.Log.Fatalf("GetPage Error: %s", err)
		os.Exit(2)
	}
	if page > 0 {
		result = (page - 1) * config.AppConfig.App.PageSize
	}

	return result
}
