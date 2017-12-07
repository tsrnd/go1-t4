package routers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/goweb4/controllers"
	ctx "github.com/astaxie/beego/context"
)

type Route struct {
	URL				    string
	Handler			  beego.ControllerInterface
	mappingMethod	[]string
}

type Routes = []Route

var routes = Routes{
	Route{
		"/",
		&controllers.MainController{},
		[]string{
			"get:Test",
		},
	},
	Route{
		"/product/add",
		&controllers.MainController{},
		[]string{
			"get:TestAdmin",
		},
	},
}

var FilterUser = func(ctx *ctx.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/login") {
		return
	}

	_, ok := ctx.Input.Session("uid").(int)
	if !ok {
			ctx.Redirect(302, "/login")
	}
}

func init() {
	if (beego.AppConfig.String("FilterUser") == "true") {
		beego.InsertFilter("/checkout", beego.BeforeRouter, FilterUser)
		beego.InsertFilter("/profile", beego.BeforeRouter, FilterUser)
	}
	for _, route := range routes {
		beego.Router(route.URL, route.Handler, route.mappingMethod...)
	}
}
