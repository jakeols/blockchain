package main

import (
	"net/http"
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
		ReceiveBlock,
	},
	Route{
		"Register",
		"GET",
		"/peer",
		Register,
	},
	Route{
		"Start",
		"GET",
		"/start",
		Start,
	},
	Route{
		"Upload",
		"GET",
		"/upload",
		Upload,
	},
	Route{
		"ReturnBlock",
		"GET",
		"/block/{height}/{hash}",
		ReturnBlock,
	},
}
