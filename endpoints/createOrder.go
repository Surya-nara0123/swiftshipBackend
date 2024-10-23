package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/helperfunction"
	"github.com/surya-nara0123/swiftship/types"
)

// CreateOrder is a function that creates a new order in the database
func CreateOrder(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	// Get the order details from the request body
	order := new(types.SafeOrder)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse JSON",
		})
	}

	// Generate a unique ID for the order
	orderId := helperfunction.GenerateUniqueInt()

	// printing the order details
	fmt.Println(*order)
	// generate unique id for order items
	itemsIDList := []int64{}
	for i := 0; i < len(order.OrderItems); i++ {
		fmt.Println("hiii")
		itemsIDList = append(itemsIDList, helperfunction.GenerateUniqueInt())
	}

	// Get the database connection
	db, _ := DbInterface.GetDbData()

	db.SavePoint("create_order")
	// Create a new order in the database
	newOrder := &types.OrderList{
		UID:           orderId,
		UserId:        order.UserId,
		RestaurantID:  order.RestuarantID,
		IsPaid:        true,
		IsCash:        true,
		TimeCreated:   order.TimeCreated,
		OrderStatusId: 2,
	}

	err := db.Create(newOrder).Error
	if err != nil {
		fmt.Println("Error: ", err.Error())
		db.RollbackTo("save_point")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create new order items for every order in the database
	for i := 0; i < len(order.OrderItems); i++ {

		foodId := order.OrderItems[i].Item
		//check if the food item exists
		foodItem := types.FoodItems{}
		db.First(&foodItem, "item = ?", foodId)
		if foodItem.UID == 0 {
			db.RollbackTo("save_point")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Food item does not exist",
			})
		}

		newOrderItem := &types.OrderDetails{
			UID:      itemsIDList[i],
			OrderId:  orderId,
			FoodId:   foodItem.UID,
			Quantity: order.OrderItems[i].Quantity,
		}

		err = db.Create(newOrderItem).Error
		if err != nil {
			fmt.Println("Error: ", err.Error())
			db.RollbackTo("save_point")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order added successfully",
		"order":   orderId,
	})
}
