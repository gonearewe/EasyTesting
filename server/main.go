package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/middlewares"
	"github.com/spf13/viper"
)

func init() {
	initViper()
	dao.InitDb()
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	middlewares.SetupMiddleWares(r)
	SetupRoute(r)
	if err := http.ListenAndServe(":"+viper.GetString("port"), r); err != nil {
		log.Fatal(err)
	}
}
