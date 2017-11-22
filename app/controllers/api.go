package controllers

import (
	"github.com/revel/revel"
	"net/url"
	"time"
	"strings"
	"strconv"
	"golang.org/x/net/idna"
)

type Api struct {
	*revel.Controller
}

type Params struct {
	url.Values
}

func (c Api) Index() revel.Result {
	// Show some documentation here

	var action string = c.Action
	var date string = time.Now().Format("2006")
	var title string = "Punycode.fun"
	return c.Render(action, title, date)
}

func (c Api) Encode() revel.Result {
	var paramStrings string
	c.Params.Bind(&paramStrings, "strings") // Sets the number of passwords

	data := make(map[string]interface{})

	stringArray := strings.Split(paramStrings,"\n")
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
	return c.RenderJSON(data)
	//return c.RenderXML(data)
}

func (c Api) Decode() revel.Result {
	var paramStrings string
	c.Params.Bind(&paramStrings, "strings") // Sets the number of passwords

	data := make(map[string]interface{})

	stringArray := strings.Split(paramStrings,"\n")
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
	return c.RenderJSON(data)
	//return c.RenderXML(data)
}

func boolToString(b bool) string {
	return strconv.FormatBool(b)
}

