package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetUnpaidOrders(c *fiber.Ctx, dbInterface database.DatabaseStruct) error {

	db, _ := dbInterface.GetDbData()

	orders := []types.OrderList{}

	db.Find(&orders, "is_paid != ?", "false")

	orderList := []types.Order{}
	for i := 0; i < len(orders); i++ {
		orderDetails := []types.OrderDetails{}

		db.Find(&orderDetails, "order_id = ?", orders[i].UID)

		orderDetails1 := []types.OrderItemReq{}
		for i := 0; i < len(orderDetails); i++ {
			foodId := orderDetails[i].FoodId
			food := types.FoodItems{}
			db.First(&food, "uid = ?", foodId)
			orderItem := types.OrderItemReq{
				Item:     food.Item,
				Quantity: orderDetails[i].Quantity,
			}
			orderDetails1 = append(orderDetails1, orderItem)
		}

		order := types.Order{
			UserId:        orders[i].UserId,
			RestuarantID:  orders[i].RestuarantID,
			IsPaid:        orders[i].IsPaid,
			IsCash:        orders[i].IsCash,
			TimeCreated:   orders[i].TimeCreated,
			OrderStatusId: orders[i].OrderStatusId,
			OrderItems:    orderDetails1,
		}

		orderList = append(orderList, order)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": orderList,
	})
}
