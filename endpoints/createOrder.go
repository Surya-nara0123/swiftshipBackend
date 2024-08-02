package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/helperfunction"
	"github.com/surya-nara0123/swiftship/types"
)

// CreateOrder is a function that creates a new order in the database
func CreateOrder(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	// Get the order details from the request body
	order := new(types.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse JSON",
		})
	}

	// Generate a unique ID for the order
	orderId := helperfunction.GenerateUniqueInt()

	// generate unique id for order items
	itemsIDList := []int32{}
	for i := 0; i < len(order.OrderItems); i++ {
		itemsIDList[i] = helperfunction.GenerateUniqueInt()
	}

	// Get the database connection
	db, _ := DbInterface.GetDbData()

	// Add the order to the database
	query := `INSERT INTO order_list (uid, user_id, restaurant_id, is_paid, is_cash, order_status_id) VALUES ($1, $2, $3, $4, $5, $6)`
	query1 := `INSERT INTO order_details (uid, order_id, food_id, quantity) VALUES ($1, $2, $3, $4)`
	_, err := db.Query(query, orderId, order.UserID, order.RestID, order.IsPaid, order.IsCash, order.OrderStatus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	for i := 0; i < len(order.OrderItems); i++ {
		_, err = db.Query(query1, itemsIDList[i], order.OrderItems[i].OrderID, order.OrderItems[i].FoodID, order.OrderItems[i].Quantity)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order added successfully",
	})
}
