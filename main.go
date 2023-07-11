package main

import (
	"go-crud-api-NoSQL/app"
)

func main() {
	// setup and run app
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
