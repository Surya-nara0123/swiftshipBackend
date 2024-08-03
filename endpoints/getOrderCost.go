package endpoints

import (
	"fmt"

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

	query := `SELECT * FROM order_list WHERE uid = $1`
	rows, err := db.Query(query, costBody.OrderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get order",
		})
	}

	query1 := `SELECT * FROM order_details WHERE order_id = $1`
	rows1, err := db.Query(query1, costBody.OrderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	order := types.OrderGet{}
	orderItems := []types.OrderItems{}

	for rows.Next() {
		err = rows.Scan(&order.ID, &order.UserID, &order.RestID, &order.IsPaid, &order.IsCash, &order.Time, &order.OrderStatus)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	for rows1.Next() {
		orderItem := types.OrderItems{}
		err = rows1.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.FoodID, &orderItem.Quantity)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		orderItems = append(orderItems, orderItem)
	}

	order.OrderItems = orderItems

	query = `SELECT * FROM food_items`
	rows, err = db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

	cost := 0

	for _, orderItem := range order.OrderItems {
		for _, foodItem := range foodItems {
			if orderItem.FoodID == foodItem.ID {
				cost += foodItem.Price * orderItem.Quantity
			}
		}
	}

	return c.JSON(fiber.Map{
		"cost": cost,
	})
}
