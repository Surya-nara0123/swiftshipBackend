package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetUnpaidOrders(c *fiber.Ctx, dbInterface database.DatabaseStruct) error {

	db, _ := dbInterface.GetDbData()

	query := `SELECT * FROM order_list WHERE is_paid = false`
	rows, err := db.Query(query)
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
