package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/handlers"
	"github.com/gonearewe/EasyTesting/middlewares"
)

func init() {
	initViper()
	writer := initLogWriter()
	initZeroLog(writer)
	handlers.InitTaskConsumers()
	dao.InitDb(writer)
	initGin(writer)
}

func main() {
	r := gin.New()
	teacherAuthRouter, adminAuthRouter, studentAuthRouter := middlewares.SetupMiddleWares(r)
	SetupRoute(r, teacherAuthRouter, adminAuthRouter, studentAuthRouter)
	port := viper.GetString("port")
	fmt.Println("Server is Running at port " + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
