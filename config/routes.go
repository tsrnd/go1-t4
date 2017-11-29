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
		"/",
		handlers.Index,
	},
	Route{
		"AddUser",
		"POST",
		"/",
		handlers.AddUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/",
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
