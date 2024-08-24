package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/endpoints"
	_ "github.com/surya-nara0123/swiftship/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://swiftship-nine.vercel.app, http://localhost:3000", // Adjust to match your frontend
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",              // Include any custom headers you use
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,OPTIONS",                         // Ensure all necessary methods are allowed
		AllowCredentials: true,                                                       // Enable if your frontend uses cookies or authentication tokens
	}))

	// Define routes
	app.Get("/readiness", handlerReadiness)

	// user endpoints
	app.Post("/createuser", func(c *fiber.Ctx) error {
		return endpoints.CreateUser(c, DbInterface)
	})
	app.Post("/getuserbyid", func(c *fiber.Ctx) error {
		return endpoints.GetUserbyID(c, DbInterface)
	})
	app.Post("/getuserid", func(c *fiber.Ctx) error {
		return endpoints.GetUserID(c, DbInterface)
	})
	app.Post("/getuserbyusername", func(c *fiber.Ctx) error {
		return endpoints.GetUserbyUsername(c, DbInterface)
	})
	app.Post("/updateUser", func(c *fiber.Ctx) error {
		return endpoints.UpdateUser(c, DbInterface)
	})

	// restaurant endpoints
	app.Post("/addrestaurant", func(c *fiber.Ctx) error {
		return endpoints.AddRestaurant(c, DbInterface)
	})
	app.Post("/getrestaurantbyid", func(c *fiber.Ctx) error {
		return endpoints.GetRestaurantbyID(c, DbInterface)
	})
	app.Post("/getresturantid", func(c *fiber.Ctx) error {
		return endpoints.GetRestaurantID(c, DbInterface)
	})
	app.Post("/getrestaurantbyname", func(c *fiber.Ctx) error {
		return endpoints.GetRestaurantbyName(c, DbInterface)
	})
	app.Get("/getrestaurants", func(c *fiber.Ctx) error {
		return endpoints.GetRestaurants(c, DbInterface)
	})

	// food endpoints
	app.Post("/addfooditems", func(c *fiber.Ctx) error {
		return endpoints.AddFoodItems(c, DbInterface)
	})
	app.Get("/getFooditems", func(c *fiber.Ctx) error {
		return endpoints.GetFoodItems(c, DbInterface)
	})
	app.Post("/getfooditemsbyrestaurant", func(c *fiber.Ctx) error {
		return endpoints.GetFoodItemsByRestaurant(c, DbInterface)
	})
	// app.Get("/getfooditemsbyid", func(c *fiber.Ctx) error {
	// 	return endpoints.GetFoodItemsByID(c, DbInterface)
	// })

	// order endpoints
	app.Post("/createorder", func(c *fiber.Ctx) error {
		return endpoints.CreateOrder(c, DbInterface)
	})
	app.Post("/getorderid", func(c *fiber.Ctx) error {
		return endpoints.GetOrderbyID(c, DbInterface)
	})
	app.Post("/getordersbyrestaurant", func(c *fiber.Ctx) error {
		return endpoints.GetOrdersbyRestaurant(c, DbInterface)
	})
	app.Post("/getordersbyuser", func(c *fiber.Ctx) error {
		return endpoints.GetOrdersbyUser(c, DbInterface)
	})
	app.Get("/getactiveorders", func(c *fiber.Ctx) error {
		return endpoints.GetActiveOrders(c, DbInterface)
	})
	app.Get("/getunpaidorders", func(c *fiber.Ctx) error {
		return endpoints.GetUnpaidOrders(c, DbInterface)
	})
	app.Post("/getordercost", func(c *fiber.Ctx) error {
		return endpoints.GetOrderCost(c, DbInterface)
	})
	app.Post("/updateorderstatus", func(c *fiber.Ctx) error {
		return endpoints.UpdateOrderStatus(c, DbInterface)
	})

	app.Get("/getcompletedorders", func(c *fiber.Ctx) error {
		return endpoints.GetCompletedOrders(c, DbInterface)
	})
	// app.Get("/getcancelledorders", func(c *fiber.Ctx) error {
	// 	return endpoints.GetCancelledOrders(c, DbInterface)
	// })

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatalln(app.Listen(fmt.Sprintf("0.0.0.0:%v", port)))
}
