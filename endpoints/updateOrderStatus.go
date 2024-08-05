package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func UpdateOrderStatus(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	orderStatus := new(types.OrderStatus)

	err := c.BodyParser(orderStatus)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	db, _ := DbInterface.GetDbData()

	_, err = db.Exec("UPDATE order_list SET order_status_id = $1 WHERE uid = $2", orderStatus.Status, orderStatus.OrderID)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println("Order status updated!")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
