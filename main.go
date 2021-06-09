package main

import (
	"fmt"
	"github.com/ehsaniara/go-crash/config"
	"github.com/ehsaniara/go-crash/models"
	"github.com/ehsaniara/go-crash/pkg/log"
	"github.com/ehsaniara/go-crash/pkg/redis"
	"github.com/ehsaniara/go-crash/routers"
	"github.com/ehsaniara/go-crash/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	config.Setup()
	log.Setup()
	models.Setup()
	redis.Setup()
	util.Setup()
}

func main() {

	// gin setting
	gin.SetMode(config.AppConfig.App.GinRunMode) //"debug","release","test"
	routersInit := routers.InitRouter()
	readTimeout := config.AppConfig.App.ReadTimeout
	writeTimeout := config.AppConfig.App.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.AppConfig.App.EndPointPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Log.Infof("start http server listening %v\n", endPoint)

	if err := server.ListenAndServe(); err != nil {
		log.Log.Errorf("failed to run the gin: %v\n", err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
