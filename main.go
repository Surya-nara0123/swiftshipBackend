package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/surya-nara0123/swiftship/database"
	_ "github.com/surya-nara0123/swiftship/types"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/endpoints"
)

func main() {
	// Initialize and open the database connection
	DbInterface := database.DatabaseStruct{} // Ensure `DatabaseStruct` is correctly defined
	_, err := DbInterface.OpenConn()
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}
	defer DbInterface.CloseConn()

	// Load environment variables
	godotenv.Load()
	app := fiber.New()

	// Define routes
	app.Get("/readiness", handlerReadiness)
	app.Post("/createuser", func(c *fiber.Ctx) error {
		return endpoints.CreateUser(c, DbInterface)
	})
	app.Get("/getuserbyid", func(c *fiber.Ctx) error {
		return endpoints.GetUserbyID(c, DbInterface)
	})
	app.Post("/addrestaurant", func(c *fiber.Ctx) error {
		return endpoints.AddRestaurant(c, DbInterface)
	})
	app.Get("/getrestaurant", func(c *fiber.Ctx) error {
		return endpoints.GetRestaurant(c, DbInterface)
	})
	app.Post("/createorder", func(c *fiber.Ctx) error {
		return endpoints.CreateOrder(c, DbInterface)
	})
	app.Get("/getorderid", func(c *fiber.Ctx) error {
		return endpoints.GetOrderbyID(c, DbInterface)
	})
	app.Get("/getuserbyusername", func(c *fiber.Ctx) error {
		return endpoints.GetUserbyUsername(c, DbInterface)
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
