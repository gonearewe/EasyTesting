package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"github.com/spf13/viper"

	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/handlers"
	"github.com/gonearewe/EasyTesting/middlewares"
)

func init() {
	initViper()
	logger.Init("EasyTesting", true, false, os.Stdout)
	handlers.InitTaskConsumers()
	dao.InitDb()
}

func main() {
	if viper.GetBool("disable_console_color") {
		gin.DisableConsoleColor()
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	teacherAuthRouter, adminAuthRouter, studentAuthRouter := middlewares.SetupMiddleWares(r)
	SetupRoute(r, teacherAuthRouter, adminAuthRouter, studentAuthRouter)
	port := viper.GetString("port")
	fmt.Println("Server is Running at port " + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
