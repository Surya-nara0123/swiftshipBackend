package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func UpdateOrderStatus(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	orderStatus := new(types.OrderStatusReq)

	err := c.BodyParser(orderStatus)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	db, _ := DbInterface.GetDbData()

	order := types.OrderList{}

	db.First(&order, "uid = ?", orderStatus.OrderID)
	if order.UID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	order.OrderStatusId = orderStatus.Status
	db.Save(&order)

	fmt.Println("Order status updated!")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
