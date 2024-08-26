package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

// EditFoodItem edits a food item in the database
func EditFoodItem(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	// Parse the request body
	food := new(types.EditFoodReq)
	if err := c.BodyParser(food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse request body",
		})
	}

	db, _ := DbInterface.GetDbData()

	foodItem := types.FoodItems{}
	db.First(&foodItem, "uid = ?", food.UID)
	foodItem.Item = food.Name
	foodItem.Price = food.Price
	foodItem.IsVeg = food.IsVeg
	foodItem.IsRegular = food.IsRegular
	foodItem.Ingredients = food.Ingredients
	foodItem.AvailableTime = food.AvailableTime

	db.Save(&foodItem)

	return c.JSON(fiber.Map{
		"status":    "ok",
		"food_item": foodItem,
	})
}
