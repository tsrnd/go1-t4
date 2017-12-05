package config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/goweb4/handlers"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/index",
		handlers.Index,
	},
	Route{
		"Register",
		"GET",
		"/register",
		handlers.Register,
	},
	Route{
		"Register",
		"POST",
		"/registerHandler",
		handlers.RegisterHandler,
	},
	Route{
		"AddUser",
		"GET",
		"/user/add",
		handlers.AddUser,
	},
	Route{
		"CreateUser",
		"POST",
		"/user",
		handlers.AddUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/user/{id}",
		handlers.UpdateUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/{id}",
		handlers.DeleteUser,
	},
	Route{
		"AddOrder",
		"GET",
		"/order/add",
		handlers.AddOrder,
	},
	Route{
		"CreateOrder",
		"POST",
		"/order",
		handlers.AddOrder,
	},
	Route{
		"EditOrder",
		"GET",
		"/order/{id}/edit",
		handlers.EditOrder,
	},
	Route{
		"UpdateOrder",
		"PUT",
		"/order/{id}",
		handlers.UpdateOrder,
	},
	Route{
		"Login",
		"GET",
		"/login",
		handlers.Login,
	},
	Route{
		"LoginHandler",
		"POST",
		"/loginHandler",
		handlers.LoginHandler,
	},
	Route{
		"LogoutHandler",
		"GET",
		"/logout",
		handlers.LogoutHandler,
	},
	Route{
		"CreateProduct",
		"GET",
		"/product/add",
		handlers.CreateProduct,
	},
	Route{
		"StoreProduct",
		"POST",
		"/product",
		handlers.StoreProduct,
	},
	Route{
		"EditProduct",
		"GET",
		"/product/{id}/edit",
		handlers.EditProduct,
	},
	Route{
		"UpdateProduct",
		"PUT",
		"/product/{id}",
		handlers.UpdateProduct,
	},
	Route{
		"DeleteProduct",
		"DELETE",
		"/product/{id}",
		handlers.DestroyProduct,
	},
	Route{
		"ShowProduct",
		"GET",
		"/product/{id}",
		handlers.ShowProduct,
	},
	Route{
		"DetailProduct",
		"GET",
		"/detail/{id}",
		handlers.DetailProduct,
	},
	Route{
		"ShowProductGroup",
		"GET",
		"/group/{id}",
		handlers.ShowProductGroup,
	},
	Route{
		"ShowCart",
		"GET",
		"/cart",
		handlers.ShowCart,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
