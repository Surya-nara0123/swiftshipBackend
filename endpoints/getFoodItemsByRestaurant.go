package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetFoodItemsByRestaurant(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	foodItem := new(types.FoodItemGet)

	err := c.BodyParser(foodItem)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	db, _ := DbInterface.GetDbData()

	rows, err := db.Query("SELECT * FROM food_items WHERE restuarant_id = $1", foodItem.ID)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	foodItems := []types.FoodItemGetResp{}
	for rows.Next() {
		foodItem := types.FoodItemGetResp{}
		err = rows.Scan(&foodItem.ID, &foodItem.RestID, &foodItem.Name, &foodItem.Ingredients, &foodItem.IsVeg, &foodItem.IsRegular, &foodItem.Price)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		foodItems = append(foodItems, foodItem)
	}

	return c.JSON(fiber.Map{
		"food_items": foodItems,
	})
}
