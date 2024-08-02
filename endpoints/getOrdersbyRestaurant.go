package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetOrdersbyRestaurant(c *fiber.Ctx, dbInterface database.DatabaseStruct) error {
	order := new(types.OrderGetRestaurant)

	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse JSON",
		})
	}

	db, _ := dbInterface.GetDbData()

	query := `SELECT * FROM order_list WHERE restaurant_id = $1`
	rows, err := db.Query(query, order.RestID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	orderList := []types.OrderGet{}

	for rows.Next() {
		order := types.OrderGet{}
		err := rows.Scan(&order.ID, &order.UserID, &order.RestID, &order.IsPaid, &order.IsCash, &order.Time, &order.OrderStatus)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		orderItemsQuery := `SELECT * FROM order_details WHERE order_id = $1`
		orderItemsRows, err := db.Query(orderItemsQuery, order.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		orderItems := []types.OrderItems{}

		for orderItemsRows.Next() {
			orderItem := types.OrderItems{}
			err := orderItemsRows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.FoodID, &orderItem.Quantity)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			orderItems = append(orderItems, orderItem)
		}

		order.OrderItems = orderItems
		orderList = append(orderList, order)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": orderList,
	})

}
