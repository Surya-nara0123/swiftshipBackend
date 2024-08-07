package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetRestaurantID(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	restaurant := new(types.RestaurantGetName)

	err := c.BodyParser(restaurant)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*restaurant)

	db, _ := DbInterface.GetDbData()

	var uid int
	query := `
	SELECT 
		uid
	FROM 
		restaurant_data 
	WHERE 
		name = $1`

	row := db.QueryRow(query, restaurant.Name)
	err = row.Scan(&uid)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println(uid)
	return c.JSON(fiber.Map{
		"uid": uid,
	})
}
