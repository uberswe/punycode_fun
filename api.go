package punycode

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/idna"
	"log"
	"net/http"
	"strings"
)

type encodeForm struct {
	Strings    string  `form:"strings" binding:"required"`
}

func encode(c *gin.Context) {
	var form encodeForm
	err := c.Bind(&form)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Invalid form data")
		return
	}

	data := make(map[string]interface{})

	stringArray := strings.Split(form.Strings,"\n")
	var resArray []string
	limit := 0
	for _, str := range stringArray {
		hash, _ := idna.ToUnicode(str)
		resArray = append(resArray, hash)
		limit++
		// Limit to 20 hashes
		if limit >= 20 {
			break
		}
	}

	data["href"] = "https://punycode.fun" + "/api/v1/encode"
	data["encoded_strings"] = resArray
	data["count"] = len(resArray)
	c.JSON(http.StatusOK, data)
}

func decode(c *gin.Context) {
	var form encodeForm
	err := c.Bind(&form)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Invalid form data")
		return
	}
	data := make(map[string]interface{})

	stringArray := strings.Split(form.Strings,"\n")
	var resArray []string
	limit := 0
	for _, str := range stringArray {
		hash, _ := idna.ToASCII(str)
		resArray = append(resArray, hash)
		limit++
		// Limit to 20 hashes
		if limit >= 20 {
			break
		}
	}

	data["href"] = "https://punycode.fun" + "/api/v1/decode"
	data["decoded_strings"] = resArray
	data["count"] = len(resArray)
	c.JSON(http.StatusOK, data)
}

