package routers

import (
	"net/http"
	"github.com/goweb4/models"
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
		"/products",
		&controllers.ProductsController{},
		[]string{
			"get:GetAll",
		},
	},
}

var FilterUser = func(ctx *ctx.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/login") {
		return
	}

	_, ok := ctx.Input.Session("uid").(int64)
	if !ok {
			ctx.Redirect(302, "/login")
	}
}

var FilterAdmin = func(ctx *ctx.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/admin/login") {
		return
	}

	uid, ok := ctx.Input.Session("uid").(int64)
	if !ok {
			ctx.Redirect(302, "/admin/login")
			return
	}

	if uid > 0 {
		user, err := models.GetUserById(uid); if err != nil {
			ctx.Redirect(302, "/") //redirect to 404 page
			return
		}
		if user.Role != models.ADMIN_ROLE {
			ctx.Redirect(http.StatusSeeOther, "/")
		}
	}
}

func init() {
	if (beego.AppConfig.String("FilterUser") == "true") {
		beego.InsertFilter("/checkout", beego.BeforeRouter, FilterUser)
		beego.InsertFilter("/profile", beego.BeforeRouter, FilterUser)
		beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterAdmin)
	}
	beego.Include(
		&controllers.UserController{},
		&controllers.ProductsController{},
		&controllers.OrderProductsController{},
		&controllers.OrdersController{},
		&controllers.ExtendController{},
		&controllers.PaymentsController{},
		&controllers.ProductGroupsController{},
		&controllers.AuthenController{},
	)
	for _, route := range routes {
		beego.Router(route.URL, route.Handler, route.mappingMethod...)
	}
}
