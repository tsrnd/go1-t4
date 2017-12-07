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
			"get:Get",
		},
	},
	Route{
		"/user",
		&controllers.UsersController{},
		[]string{
			"get:GetAll",
		},
	},
	Route{
		"/product_groups",
		&controllers.ProductGroupsController{},
		[]string{
			"get:GetAll",
		},
	},
	Route{
		"/products",
		&controllers.ProductsController{},
		[]string{
			"get:GetAll",
		},
	},
}

func init() {
	for _, route := range routes {
		beego.Router(route.URL, route.Handler, route.mappingMethod...)
	}
}
