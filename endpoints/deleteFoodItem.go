package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

// DeleteFoodItem deletes a food item from the database
func DeleteFoodItem(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	// Parse the request body
	type Request struct {
		FoodItemID int `json:"foodItemID"`
	}
	request := Request{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse request body",
		})
	}

	db, _ := DbInterface.GetDbData()

	db.Delete(&types.FoodItems{}, "uid = ?", request.FoodItemID)

	return c.SendStatus(fiber.StatusOK)
}
