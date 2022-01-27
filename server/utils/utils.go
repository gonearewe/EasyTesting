package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
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

func IsAnagram(str1 string,str2 string)bool{
	array1 := []rune(str1)
	array2 := []rune(str2)
	sort.Slice(array1, func(i int, j int) bool {
		return array1[i] < array1[j]
	})
	sort.Slice(array2, func(i int, j int) bool {
		return array2[i] < array2[j]
	})
	for i := range array1 {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}
