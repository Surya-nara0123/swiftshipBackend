package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/helperfunction"
	"github.com/surya-nara0123/swiftship/types"
)

func AddFoodItems(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	foodItem := new(types.FoodItem)

	err := c.BodyParser(foodItem)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*foodItem)

	//generate unique id
	uid := helperfunction.GenerateUniqueInt()
	fmt.Println(uid)

	db, _ := DbInterface.GetDbData()

	// fmt.Println(foodItem.Ingredients)

	_, err = db.Exec("INSERT INTO food_items (uid, restuarant_id, item, ingredients, is_veg, is_regular, price) VALUES ($1, $2, $3, $4, $5, $6, $7)", uid, foodItem.RestID, foodItem.Name, foodItem.Ingredients, foodItem.IsVeg, foodItem.IsRegular, foodItem.Price)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println("Food item created!")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
