package punycode

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type baseData struct {
	Title string
	MoreStyles []string
	MoreScripts []string
	Action string
	Date string
	StringsVarEncoded string
	StringsVarDecoded string
	Remember string
}

func index(c *gin.Context) {
	var result baseData
	result.Date = time.Now().Format("2006")
	result.Title = "Punycode.fun"

	s := sessions.Default(c)
	stringsVarDecoded := "sök.xyz"

	tmpStrings, found := c.Params.Get("stringss")
	if found {
		stringsVarDecoded = tmpStrings
	}

	if s.Get("strings") != nil && s.Get("strings") != "" {
		stringsVarDecoded = s.Get("strings").(string)
	}

	remember := "false"
	if s.Get("remember") != nil {
		remember = s.Get("remember").(string)
	}

	result.StringsVarDecoded = stringsVarDecoded
	result.Remember = remember

	// Should be moved to new controller and all controllers inherit
	result.Action = "App.Index"
	c.HTML(http.StatusOK, "index.html", result)
}
