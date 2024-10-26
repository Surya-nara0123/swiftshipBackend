package endpoints

// parameters are cart, userId,
import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	razorpay "github.com/razorpay/razorpay-go"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func CallRazorPay(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	// Get the order details from the request body
	type checkoutData struct {
		UserId     int64                `json:"user_id"`
		OrderItems []types.OrderItemReq `json:"order_items"`
	}
	order := new(checkoutData)
	if err := c.BodyParser(order); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse JSON",
		})
	}

	db, _ := DbInterface.GetDbData()

	// calculate the amount for the order
	amount := 0
	fmt.Println(order.OrderItems)
	for i := 0; i < len(order.OrderItems); i++ {
		// using the food name to get the price of the food
		food := &types.FoodItems{}
		db.First(&food, "item = ?", order.OrderItems[i].Item)
		fmt.Println(order.OrderItems[i].Item)
		amount += food.Price * order.OrderItems[i].Quantity
	}

	fmt.Println(amount)
	razorpayClient := razorpay.NewClient(os.Getenv("RAZORPAY_KEY"), os.Getenv("RAZORPAY_SECRET"))
	result, err := razorpayClient.Order.Create(
		map[string]interface{}{
			"amount":          amount * 100,
			"currency":        "INR",
			"receipt":         "receipt#1",
			"partial_payment": false,
		},
		nil,
	)
	fmt.Println(result, err)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"options": result,
		"key":     os.Getenv("RAZORPAY_KEY"),
	})
}
