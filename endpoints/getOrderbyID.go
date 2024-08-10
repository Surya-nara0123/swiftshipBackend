package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetOrderbyID(c *fiber.Ctx, dbInterface database.DatabaseStruct) error {

	orderIdStruct := new(types.OrderIDReq)

	err := c.BodyParser(orderIdStruct)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	db, _ := dbInterface.GetDbData()

	orderData := types.OrderList{}

	db.First(&orderData, "uid = ?", orderIdStruct.ID)

	orderDetails := []types.OrderDetails{}

	db.Find(&orderDetails, "order_id = ?", orderData.UID)

	/*
		type Order struct {
			UserId        int64  `json:"user_id"`
			RestuarantID  int64  `json:"rest_id"`
			IsPaid        bool   `json:"is_paid"`
			IsCash        bool   `json:"is_cash"`
			TimeCreated   string `json:"timestamp"`
			OrderStatusId int    `json:"order_status"`
			OrderItems    []OrderItemReq
		}
	*/
	order := types.Order{
		UserId:        orderData.UserId,
		RestuarantID:  orderData.RestuarantID,
		IsPaid:        orderData.IsPaid,
		IsCash:        orderData.IsCash,
		TimeCreated:   orderData.TimeCreated,
		OrderStatusId: orderData.OrderStatusId,
		OrderItems:    []types.OrderItemReq{},
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"order": order,
	})
}
