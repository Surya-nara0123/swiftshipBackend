package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

// GetRestaurants is a function that returns all the restaurants in the database
func GetRestaurants(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	db, _ := DbInterface.GetDbData()

	restaurantData := []types.RestaurantData{}

	db.Find(&restaurantData)

	return c.JSON(fiber.Map{
		"result": restaurantData,
		"status": "ok",
	})
}
