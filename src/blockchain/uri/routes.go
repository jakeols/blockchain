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
		"SimpleGet",
		"GET",
		"/simpleget",
		handlers.SimpleGet,
	},
	Route{
		"AnotherGet",
		"GET",
		"/anotherget/{number}",
		handlers.AnotherGet,
	},
	Route{
		"SimplePost",
		"POST",
		"/simplepost",
		handlers.SimplePost,
	},
	Route{
		"AskOddOrEven",
		"GET",
		"/askoddoreven/{number}", // sample => askoddoreven/5
		handlers.AskOddOrEven,    //api to send post request
	},
	Route{
		"OddOrEven",
		"POST",
		"/oddoreven",
		handlers.OddOrEven,
	},
}
