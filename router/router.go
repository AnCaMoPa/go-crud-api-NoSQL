package router

import (
	"go-crud-api-NoSQL/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Check if the server is up and running
	app.Get("/health", handlers.HandleHealthCheck)

	// setup the crud group in which we can find all the method
	crud := app.Group("/crud")

	// Declare the diferent methods in the crud group

	//Params
	//:collection = Name of the DB collection you want the end point to work with
	//:id = MongoDB ID of the Document inside the collection you want the endpoint to work with

	//Get endpoints
	crud.Get("/:collection", handlers.HandleAllCollections)       //Get all the documents of a collection
	crud.Get("/:collection/:id", handlers.HandleGetOneCollection) //Get the document with the said id of the collection

	//Post endpoint
	crud.Post("/:collection", handlers.HandleCreateCollection) // Create a new document in the specified collection

	//Put endpoints
	crud.Put("/:collection/:id", handlers.HandleUpdateCollection) //Update the document with the said id of the collection
	crud.Put("/:collection", handlers.HandleAllUpdateCollection)  //Update all documents of the collection

	//Delete endpoints
	crud.Delete("/:collection/:id", handlers.HandleDeleteCollection) //Delete the document with the said id of the collection
	crud.Delete("/:collection", handlers.HandleAllDeleteCollection)  //Delete all documents of the collection

}
