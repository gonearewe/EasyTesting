package handlers

import (
    "io/ioutil"
    "strconv"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/utils"
    "github.com/google/logger"
    "github.com/patrickmn/go-cache"
)


func GetCacheHandler(c *gin.Context) {
    key := strconv.Itoa(int(jwt.ExtractClaims(c)["exam_session_id"].(float64)))
    if ret,found := utils.MemoryStore.Get(key);found{
        c.String(200, ret.(string))
    }else {
        logger.Info("no cache found for exam_session_id: "+key)
    }
}

func PutCacheHandler(c *gin.Context) {
    key := strconv.Itoa(int(jwt.ExtractClaims(c)["exam_session_id"].(float64)))
    var body, err = ioutil.ReadAll(c.Request.Body)
    utils.PanicWhen(err)
    s := string(body)
    utils.MemoryStore.Set(key, s,cache.DefaultExpiration)
}
