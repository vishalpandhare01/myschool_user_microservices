package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

	initializer.ConnectDb()
	initializer.Migration()
	internal.RouteSetUp(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
