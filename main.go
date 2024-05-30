package main

import (
    "os"
    "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rafia9005/GoLuva/database"
	"github.com/rafia9005/GoLuva/routes"
    "github.com/joho/godotenv"
)

func main() {
    // load dotnev
    godotenv.Load()
    
    // create gofiber instance
	app := fiber.New()

    var appPort = os.Getenv("APP_PORT")
    if appPort == "" {
        appPort = "8000"
    }

    if(os.Getenv("APP_DATABASE") != ""){
        database.Connect()
        routes.AutoMigrate()
    }

	routes.SetupRouter(app)

    app.Listen(fmt.Sprintf(":%s", appPort))
}
