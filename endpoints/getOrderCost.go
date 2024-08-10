package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetOrderCost(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	costBody := new(types.CostBody)

	err := c.BodyParser(costBody)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	db, _ := DbInterface.GetDbData()

	orderDetails := []types.OrderDetails{}

	db.Find(&orderDetails, "order_id = ?", costBody.OrderID)

	cost := 0
	for i := 0; i < len(orderDetails); i++ {
		foodID := orderDetails[i].FoodId
		food := types.FoodItems{}
		db.First(&food, "uid = ?", foodID)
		cost += food.Price * orderDetails[i].Quantity
	}

	return c.JSON(fiber.Map{
		"cost": cost,
	})
}
