package main

import (
	"errors"
	"fmt"

	// "github.com/gofiber/fiber/v2"
	"configcat-homework/app/controllers"
	"configcat-homework/app/middleware"
	"configcat-homework/app/routing"

	// "configcat-homework/internal/services"

	"configcat-homework/internal/custom_configcat"
	"configcat-homework/internal/custom_viper"
	// configcatAbstract "configcat-homework/internal/configcat"

	// configcat "github.com/configcat/go-sdk"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/spf13/viper"
)

func main() {
	appViper := custom_viper.Set{}.Viper().CustomViper

	
	env, ok := appViper.Get("ENVIRONMENT").(string)

	if !ok {
		err := errors.New("ENVIRONMENT not found")
		fmt.Println(err)
	}

	fmt.Println("Starting server via", "environment", env)

	fiberApp := bootstrapFiber()

	startServer(fiberApp, appViper)
}


// Setup Setup a fiber app with all of its routes
func bootstrapFiber() *fiber.App {
	// Configure cors
	var corsConfig = cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
	}

	// Setup error handling
	fiberConfig := fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	}

	loggerConfig := fiberLogger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}

	// Initialize a new app
	app := fiber.New(fiberConfig)
	app.Use(fiberLogger.New(loggerConfig))
	app.Use(fiberRecover.New())
	app.Use(cors.New(corsConfig))

	controllerManager := controllers.NewControllerManager()

	// Setup routing
	baseRouter := app.Group("/api/v1")
	routing.CreatePublicRoutes(baseRouter, controllerManager)

	// Return the configured app
	return app
}


func startServer(app *fiber.App, appViper *viper.Viper) {
	port, ok := appViper.Get("PORT").(string)

	if !ok {
		err := errors.New("PORT not found")
		fmt.Println(err)
	}

	listenErr := app.Listen(":" + port)

	if listenErr != nil {
		fmt.Println("Error during listen", listenErr)
	}
}