package controllers

import (
	"github.com/revel/revel"
	"time"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var date string = time.Now().Format("2006")
	var title string = "Punycode.fun"
	var count string = "5"

	stringsVarDecoded := "sök.xyz"

	var stringsTempVar string
	c.Params.Query=c.Request.URL.Query()
	c.Params.Bind(&stringsTempVar,"strings")

	if (c.Session["strings"] != "") {
		stringsVarDecoded = c.Session["strings"]
	}

	if (stringsTempVar != "") {
		stringsVarDecoded = stringsTempVar
	}

	var remember string = "false"
	if (c.Session["remember"] != "") {
		remember = c.Session["remember"];
	}

	// Should be moved to new controller and all controllers inherit
	var action string = c.Action
	return c.Render(date, title, count, stringsVarDecoded, remember, action)
}
