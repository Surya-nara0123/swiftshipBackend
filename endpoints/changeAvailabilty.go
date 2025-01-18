package endpoints

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func ChangeAvailability(c *fiber.Ctx, dbInterface database.DatabaseStruct) error {
	foodId := new(types.RestaurantAvailability)
	err := c.BodyParser(foodId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	fmt.Println(foodId)
	
	db, _ := dbInterface.GetDbData()
	
	foodItem := types.FoodItems{}
	db.First(&foodItem, "uid = ?", foodId.FoodId)
	
	fmt.Println(foodId)
	foodItem.IsAvailable = !foodItem.IsAvailable

	db.Where("uid = ?", foodId.FoodId).Save(&foodItem)

	return c.JSON(fiber.Map{})

}
