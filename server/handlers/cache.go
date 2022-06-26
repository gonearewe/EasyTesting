package handlers

import (
	"io/ioutil"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"

	"github.com/gonearewe/EasyTesting/utils"
)

func GetCacheHandler(c *gin.Context) {
	key := strconv.Itoa(int(jwt.ExtractClaims(c)["exam_session_id"].(float64)))
	if ret, found := utils.MemoryStore.Get(key); found {
		c.String(200, ret.(string))
	} else {
		log.Info().Str("exam_session_id", key).Msg("no cache found")
	}
}

func PutCacheHandler(c *gin.Context) {
	key := strconv.Itoa(int(jwt.ExtractClaims(c)["exam_session_id"].(float64)))
	var body, err = ioutil.ReadAll(c.Request.Body)
	utils.PanicWhen(err)
	s := string(body)
	utils.MemoryStore.Set(key, s, cache.DefaultExpiration)
}
