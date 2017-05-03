package main

import (
	"os"

	"github.com/timkellogg/ok-go/config"
	"github.com/timkellogg/ok-go/routes"
)

func main() {
	app := config.NewApp(".")
	routes.Setup(app)
	app.Negroni.Run(":" + os.Getenv("PORT"))
}
