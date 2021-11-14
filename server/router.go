package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/handlers"
	"github.com/gonearewe/EasyTesting/utils"
	"github.com/spf13/viper"
)

func SetupAuth(r gin.IRoutes) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "easy testing",
		Key:           []byte(viper.GetString("jwt_secret_key")),
		IdentityKey:   "id",
		Authenticator: handlers.TeacherAuthenticator,
		// PayloadFunc: func(data interface{}) jwt.MapClaims {
		//     if v, ok := data.(*User); ok {
		//         return jwt.MapClaims{
		//             identityKey: v.UserName,
		//         }
		//     }
		//     return jwt.MapClaims{}
		// },
		// IdentityHandler: func(c *gin.Context) interface{} {
		//     claims := jwt.ExtractClaims(c)
		//     return &User{
		//         UserName: claims[identityKey].(string),
		//     }
		// },

		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
	})

	utils.PanicWhen(err)

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/teachers", handlers.TeachersRegisterHandler)
}
