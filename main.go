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
		return endpoints.GetRestaurantbyID(c, DbInterface)
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
	app.Get("/getrestaurantbyname", func(c *fiber.Ctx) error {
		return endpoints.GetRestaurantbyName(c, DbInterface)
	})
	app.Get("/getordersbyrestaurant", func(c *fiber.Ctx) error {
		return endpoints.GetOrdersbyRestaurant(c, DbInterface)
	})
	app.Get("/getordersbyuser", func(c *fiber.Ctx) error {
		return endpoints.GetOrdersbyUser(c, DbInterface)
	})
	app.Get("/getactiveorders", func(c *fiber.Ctx) error {
		return endpoints.GetActiveOrders(c, DbInterface)
	})
	app.Get("/getunpaidorders", func(c *fiber.Ctx) error {
		return endpoints.GetUnpaidOrders(c, DbInterface)
	})
	app.Post("/addfooditems", func(c *fiber.Ctx) error {
		return endpoints.AddFoodItems(c, DbInterface)
	})
	app.Get("/getfooditemsbyrestaurant", func(c *fiber.Ctx) error {
		return endpoints.GetFoodItemsByRestaurant(c, DbInterface)
	})
	app.Post("/getordercost", func(c *fiber.Ctx) error {
		return endpoints.GetOrderCost(c, DbInterface)
	})
	app.Post("/updateorderstatus", func(c *fiber.Ctx) error {
		return endpoints.UpdateOrderStatus(c, DbInterface)
	})
	app.Get("/getresturantid", func(c *fiber.Ctx) error {
		return endpoints.GetRestaurantID(c, DbInterface)
	})
	app.Get("/getuserid", func(c *fiber.Ctx) error {
		return endpoints.GetUserID(c, DbInterface)
	})
	// app.Get("/getcompletedorders", func(c *fiber.Ctx) error {
	// 	return endpoints.GetCompletedOrders(c, DbInterface)
	// })
	// app.Get("/getcancelledorders", func(c *fiber.Ctx) error {
	// 	return endpoints.GetCancelledOrders(c, DbInterface)
	// })
	// app.Get("/getfooditemsbyid", func(c *fiber.Ctx) error {
	// 	return endpoints.GetFoodItemsByID(c, DbInterface)
	// })

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
