package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

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

// Int trys to parse `num` to an integer, otherwise panics.
func Int(num string) int {
	ret, err := strconv.Atoi(num)
	PanicWhen(err)
	return ret
}

func Join(a []interface{}) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = fmt.Sprint(v)
	}
	return strings.Join(b, "")
}

func PanicWhen(err error) {
	if err != nil {
		panic(err)
	}
}
