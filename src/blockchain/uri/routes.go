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
	Route{
		"Start",
		"GET",
		"/start",
		handlers.Start,
	},
	Route{
		"Upload",
		"GET",
		"/upload",
		handlers.Upload,
	},
	Route{
		"ReturnBlock",
		"GET",
		"/block/{height}/{hash}",
		handlers.ReturnBlock,
	},
}
