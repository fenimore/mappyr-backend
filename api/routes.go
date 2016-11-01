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
	}, /* Database Routes */
	// READ
	Route{
		"ShowComment",
		"GET",
		"/comment/{id}", // Comment id
		ShowComment,
	},
	Route{
		"ShowLocal",
		"POST",
		"/local",
		ShowLocalComments,
	},
	Route{
		"ShowComments",
		"GET",
		"/all/comments",
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
		"/upvote/{comment_id}",
		UpVote,
	},
	Route{
		"DownVote",
		"GET",
		"/downvote/{comment_id}",
		DownVote,
	},
	Route{
		"Delete",
		"GET",
		"/delete/{id}",
		DeleteComment,
	},
	/* User Routes */
	Route{
		"SignUp",
		"POST",
		"/signup",
		Signup,
	},
	Route{
		"ShowUser",
		"GET",
		"/user/{id}",
		ShowUser,
	},
	Route{
		"ShowUsers",
		"GET",
		"/all/users",
		ShowUsers,
	},
	Route{
		"UserVotes",
		"GET",
		"/votes/{user_id}",
		UserVotes,
	},
	Route{
		"UserComments",
		"GET",
		"/comments/{user_id}",
		UserComments,
	},
	// TODO: Create
	// Can't create unless username doesn't exits
	// TODO: Delete User
	/* Authentication */
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
