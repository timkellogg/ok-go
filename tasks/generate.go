package tasks

import (
	"fmt"

	"github.com/chuckpreslar/gofer"
)

// GenerateModel - creates a model file, and test file
var GenerateModel = gofer.Register(gofer.Task{
	Label:       "model",
	Description: "Generates a model and a test file",
	Action: func(arguments ...string) error {
		fmt.Println(arguments)
		if len(arguments) < 1 {
			return Error.new("Error")
		}
		return nil
	},
})

// Handles the generation of models, scaffolds, migrations, etc.
