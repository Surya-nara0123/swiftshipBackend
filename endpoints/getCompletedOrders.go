package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetCompletedOrders(c *fiber.Ctx, dbInterface database.DatabaseStruct) error {
	user := types.FoodItemUserId{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	db, _ := dbInterface.GetDbData()

	order := []types.OrderList{}

	db.Where("user_id = ? ", user.ID).Find(&order, "order_status_id = ?", "5")

	orderList := []types.OrderCustomer{}
	for i := 0; i < len(order); i++ {
		orderDetails := []types.OrderDetails{}

		db.Find(&orderDetails, "order_id = ?", order[i].UID)

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

		order := types.OrderCustomer{
			UID:           order[i].UID,
			UserId:        order[i].UserId,
			RestuarantID:  order[i].RestaurantID,
			IsPaid:        order[i].IsPaid,
			IsCash:        order[i].IsCash,
			TimeCreated:   order[i].TimeCreated,
			OrderStatusId: order[i].OrderStatusId,
			OrderItems:    orderDetails1,
		}

		orderList = append(orderList, order)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": orderList,
	})
}
