package routers

import (
	"github.com/astaxie/beego"
	"github.com/goweb4/controllers"
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

func init() {
	for _, route := range routes {
		beego.Router(route.URL, route.Handler, route.mappingMethod...)
	}
}
