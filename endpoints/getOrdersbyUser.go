package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetOrdersbyUser(c *fiber.Ctx, dbInterface database.DatabaseStruct) error {
	order := new(types.OrderUserIDReq)

	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse JSON",
		})
	}

	db, _ := dbInterface.GetDbData()

	orderList := []types.OrderList{}

	db.Find(&orderList, "rest_id = ?", order.UserID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": orderList,
	})

}
