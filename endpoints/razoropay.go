package endpoints

// parameters are cart, userId,
import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	razorpay "github.com/razorpay/razorpay-go"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
	"golang.org/x/crypto/bcrypt"
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
	for i := 0; i < len(order.OrderItems); i++ {
		// using the food name to get the price of the food
		food := &types.Food{}
		db.Where("item = ?", order.OrderItems[i].Item).First(food)
		fmt.Println(food)
		amount += food.Price * order.OrderItems[i].Quantity
	}

	fmt.Println(amount)
	razorpayClient := razorpay.NewClient(os.Getenv("RAZORPAY_KEY"), os.Getenv("RAZORPAY_SECRET"))
	result, err := razorpayClient.Order.Create(
		map[string]interface{}{
			"amount":          amount,
			"currency":        "INR",
			"receipt":         "receipt#1",
			"partial_payment": false,
		},
		nil,
	)
	fmt.Println(result, err)
	hashedAmount, err := bcrypt.GenerateFromPassword([]byte(string(rune(amount))), 123)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Could not hash password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"amount": hashedAmount,
	})
}
