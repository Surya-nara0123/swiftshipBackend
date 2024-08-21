package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/helperfunction"
	"github.com/surya-nara0123/swiftship/types"
)

func AddFoodItems(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	food := new(types.Food)

	err := c.BodyParser(food)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*food)

	//generate unique id
	uid := helperfunction.GenerateUniqueInt()
	fmt.Println(uid)

	db, _ := DbInterface.GetDbData()

	foodItem := &types.FoodItems{
		UID:          uid,
		Item:         food.Name,
		Price:        food.Price,
		IsVeg:        food.IsVeg,
		RestuarantId: food.RestID,
		IsRegular:    food.IsRegular,
	}

	err = db.Create(foodItem).Error
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
