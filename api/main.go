package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/config"
	"github.com/sg3t41/api/model"
	"github.com/sg3t41/api/pkg/redis"
	"github.com/sg3t41/api/pkg/util"
	"github.com/sg3t41/api/router"
)

func init() {
	config.Setup("config/app.ini")
	util.Setup()
	model.Setup()
	redis.SetUp()
}

func main() {
	gin.SetMode(config.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
