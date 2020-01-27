package routes

import (
	"bacon-api/handlers"
	"net/http"
)

// Route struct
type Route struct {
	Name string
	Path string
	HTTP string
	Func http.HandlerFunc
}

// Routes struct
type Routes []Route

// Router export
var Router = Routes{
	Route{
		"NewMeeting",
		"/newmeetings",
		"POST",
		handlers.NewMeeting,
	},
	Route{
		"SetMeeting",
		"/setmeetings",
		"POST",
		handlers.SetMeeting,
	},
	Route{
		"GetMeeting",
		"/{id}",
		"GET",
		handlers.GetMeeting,
	},
}
