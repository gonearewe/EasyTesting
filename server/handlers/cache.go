package handlers

import (
    "io/ioutil"
    "strconv"
    "time"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/middlewares"
    "github.com/gonearewe/EasyTesting/utils"
)

func GetCacheHandler(c *gin.Context) {
    key := strconv.Itoa(int(jwt.ExtractClaims(c)["exam_session_id"].(float64)))
    var ret string
    err := middlewares.MemoryStore.Get(key, &ret)
    utils.PanicWhen(err)
    c.String(200, ret)
}

func PutCacheHandler(c *gin.Context) {
    key := strconv.Itoa(int(jwt.ExtractClaims(c)["exam_session_id"].(float64)))
    var body, err = ioutil.ReadAll(c.Request.Body)
    utils.PanicWhen(err)
    s := string(body)
    err = middlewares.MemoryStore.Set(key, s, 3*time.Hour)
    if err != nil {
        err = middlewares.MemoryStore.Add(key, s, 3*time.Hour)
        utils.PanicWhen(err)
    }
}
