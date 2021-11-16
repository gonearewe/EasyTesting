package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/middlewares"
	"github.com/google/logger"
	"github.com/spf13/viper"
)

func init() {
	initViper()
	logger.Init("EasyTesting", true, false, os.Stdout)
	dao.InitDb()
}

func main() {
	r := gin.New()
	teacherAuthRouter, adminAuthRouter := middlewares.SetupMiddleWares(r)
	SetupRoute(r, teacherAuthRouter, adminAuthRouter)
	if err := http.ListenAndServe(":"+viper.GetString("port"), r); err != nil {
		log.Fatal(err)
	}
}
