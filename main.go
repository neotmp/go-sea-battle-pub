package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/neotmp/go-sea-battle/database"
	"github.com/neotmp/go-sea-battle/routes"
)

func main() {

	database.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     os.Getenv("ALLOW_ORIGINS"),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	routes.Setup(app)

	if os.Getenv("ENV") == "dev" {
		app.Listen(":4003")
	} else {
		log.Fatal(app.ListenTLS(":4003", "./cert.pem", "./cert.key"))
	}

}
