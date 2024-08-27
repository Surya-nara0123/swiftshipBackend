package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetOrdersbyRestaurant(c *fiber.Ctx, dbInterface database.DatabaseStruct) error {
	order := new(types.OrderRestIDReq)

	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse JSON",
		})
	}

	fmt.Println(order.RestID)

	db, _ := dbInterface.GetDbData()

	orderList := []types.OrderList{}

	db.Find(&orderList, "restaurant_id = ?", order.RestID)

	orders := []types.OrderVendor{}

	for i := 0; i < len(orderList); i++ {
		orderDetails := []types.OrderDetails{}
		db.Find(&orderDetails, "order_id = ?", orderList[i].UID)

		orderItems := []types.OrderItemReq2{}

		for j := 0; j < len(orderDetails); j++ {
			foodItem := types.FoodItems{}
			db.First(&foodItem, "uid = ?", orderDetails[j].FoodId)
			orderItems = append(orderItems, types.OrderItemReq2{
				Item:     foodItem.Item,
				Price:    foodItem.Price,
				Quantity: orderDetails[j].Quantity,
			})
		}

		user := types.UserDetails{}
		db.First(&user, "uid = ?", orderList[i].UserId)

		orders = append(orders, types.OrderVendor{
			UID:           orderList[i].UID,
			UserName:      user.Username,
			RestuarantID:  orderList[i].RestaurantID,
			IsPaid:        orderList[i].IsPaid,
			IsCash:        orderList[i].IsCash,
			TimeCreated:   orderList[i].TimeCreated,
			OrderStatusId: orderList[i].OrderStatusId,
			OrderItems:    orderItems,
		})
	}

	// if len(orders) == 0 {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"error": "No orders found",
	// 	})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": orders,
	})

}
