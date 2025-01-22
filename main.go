package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Env not loaded", err)
	}
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000/, http://localhost:3001/, http://localhost:3002/",
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
	}))

	initializer.ConnectDb()
	initializer.Migration()
	internal.RouteSetUp(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
