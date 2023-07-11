package app

import (
	"go-crud-api-NoSQL/config"
	"go-crud-api-NoSQL/database"
	"go-crud-api-NoSQL/router"
	"os" //It allow to work with the system

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupAndRunApp() error {

	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	// defer closing database, this will close the MongoDB conecction after the App finish all the operations.
	defer database.CloseMongoDB()

	// create app
	app := fiber.New()

	//Middlewares
	app.Use(cors.New())    //Cors errors, this middleware solve the Cors problems.
	app.Use(recover.New()) //This middleware will make sure that the app will recover in case of and unexpected error occur.
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	})) //This moddleware will show us log outputs.

	// setup endpoints
	router.SetupRoutes(app)

	// get the port of the server and start the App
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
