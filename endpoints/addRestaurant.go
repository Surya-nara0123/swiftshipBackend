package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/helperfunction"
	"github.com/surya-nara0123/swiftship/types"
)

func AddRestaurant(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	restaurant := new(types.Restuarant)

	err := c.BodyParser(restaurant)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*restaurant)

	//generate unique id
	uid := helperfunction.GenerateUniqueInt()

	db, _ := DbInterface.GetDbData()

	_, err = db.Exec("INSERT INTO restaurant_data (uid, name, location, is_veg) VALUES ($1, $2, $3, $4)", uid, restaurant.Name, restaurant.Location, restaurant.IsVeg)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error": "Error while creating restaurant",
		})
	}

	fmt.Println("Restaurant created!")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
