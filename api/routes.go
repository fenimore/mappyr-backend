package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

// Define handlers in handlers.go
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	}, // READ
	Route{
		"ShowComment",
		"GET",
		"/comment/{id}",
		ShowComment,
	},
	Route{
		"ShowComments",
		"GET",
		"/comments",
		ShowComments,
	}, // CREATE
	Route{
		"NewComment",
		"POST",
		"/new",
		NewComment,
	}, // UPDATE
	Route{
		"UpVote",
		"GET",
		"/upvote/{id}",
		UpVote,
	},
	Route{
		"DownVote",
		"GET",
		"/downvote/{id}",
		DownVote,
	},
	Route{
		"Delete",
		"GET",
		"/delete/{id}",
		DeleteComment,
	}, // AUTH
	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
	Route{ // For testing purposes
		"NewToken",
		"GET",
		"/token/{id}",
		NewToken,
	},
}
