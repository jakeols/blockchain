package uri

import (
	"net/http"

	"./handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ReceiveBlock",
		"POST",
		"/block/receive/",
		handlers.ReceiveBlock,
	},
	Route{
		"Register",
		"GET",
		"/peer",
		handlers.Register,
	},
}
