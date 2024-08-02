package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetOrdersbyUser(c *fiber.Ctx, dbInterface database.DatabaseStruct) error {
	orderId := new(types.GetOrder)

	if err := c.BodyParser(orderId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse JSON",
		})
	}

	db, _ := dbInterface.GetDbData()

	query := `SELECT * FROM order_list WHERE uid = $1`
	rows, err := db.Query(query, orderId.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get order",
		})
	}

	query1 := `SELECT * FROM order_details WHERE order_id = $1`
	rows1, err := db.Query(query1, orderId.ID)
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"order": order,
	})
}
