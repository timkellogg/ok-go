package routes

import (
	"github.com/gorilla/schema"
	"github.com/timkellogg/ok-go/config"
)

var App *config.App
var decoder = schema.NewDecoder()

func Setup(app *config.App) {

	App = app

	// Define your routes here:

	App.Router.HandleFunc("/", HomeIndex).Methods("GET")

	App.Router.HandleFunc("/posts", PostsIndex).Methods("GET")
	App.Router.HandleFunc("/posts/new", PostsNew).Methods("GET")
	App.Router.HandleFunc("/posts", PostsCreate).Methods("POST")
}
