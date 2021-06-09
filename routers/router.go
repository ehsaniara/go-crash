package routers

import (
	"github.com/ehsaniara/go-crash/middleware/jwt"
	"github.com/ehsaniara/go-crash/routers/api"
	"github.com/ehsaniara/go-crash/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", Ping)
	r.POST("/auth", api.GetAuth)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/customers/:id", v1.GetCustomerById)
		apiV1.POST("/customers/", v1.AddNewCustomer)
	}

	return r
}
