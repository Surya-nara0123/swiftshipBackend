package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/endpoints"
	"github.com/surya-nara0123/swiftship/middleware"
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
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-API-Key",   // Adjust to match your frontend
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,OPTIONS",                         // Ensure all necessary methods are allowed
		AllowCredentials: true,                                                       // Enable if your frontend uses cookies or authentication tokens
	}))

	// Define routes
	app.Get("/readiness", handlerReadiness)

	// user endpoints
	app.Post("/createnormaluser", func(c *fiber.Ctx) error {
		return endpoints.CreateNormalUser(c, DbInterface)
	})
	// app.Post("/createvendoruserv123", func(c *fiber.Ctx) error {
	// 	return endpoints.CreateUser(c, DbInterface)
	// })
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
	app.Post("/changeavailability", func(c *fiber.Ctx) error {
		return endpoints.ChangeAvailability(c, DbInterface)
	})
	app.Post("/deletefooditem", func(c *fiber.Ctx) error {
		return endpoints.DeleteFoodItem(c, DbInterface)
	})
	app.Post("/editfooditem", func(c *fiber.Ctx) error {
		return endpoints.EditFoodItem(c, DbInterface)
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
	app.Post("/getactiveorders", func(c *fiber.Ctx) error {
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

	app.Post("/getcompletedorders", func(c *fiber.Ctx) error {
		return endpoints.GetCompletedOrders(c, DbInterface)
	})
	app.Post("/razorpay", func(c *fiber.Ctx) error {
		return endpoints.CallRazorPay(c, DbInterface)
	})

	//cookie endpoints
	app.Get("/getcookies", func(c *fiber.Ctx) error {
		return middleware.GetCookies(c)
	})
	app.Get("/clearcookies", func(c *fiber.Ctx) error {
		return middleware.ClearCookies(c)
	})
	app.Post("/setcookie", func(c *fiber.Ctx) error {
		return middleware.SetCookie(c)
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatalln(app.Listen(fmt.Sprintf("0.0.0.0:%v", port)))
}
