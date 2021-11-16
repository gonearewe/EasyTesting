package utils

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MustParseJson(c *gin.Context) gin.H {
	var jsonData gin.H
	MustParseJsonTo(c, &jsonData)
	return jsonData
}

func MustParseJsonTo(c *gin.Context, container interface{}) {
	data, err := ioutil.ReadAll(c.Request.Body)
	PanicWhen(err)
	err = json.Unmarshal(data, container)
	PanicWhen(err)
}

func Int(num string) int {
	ret, err := strconv.Atoi(num)
	PanicWhen(err)
	return ret
}

func PanicWhen(err error) {
	if err != nil {
		panic(err)
	}
}
