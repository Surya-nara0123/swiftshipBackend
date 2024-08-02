package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetRestaurantbyID(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	restaurant := new(types.RestaurantGet)

	err := c.BodyParser(restaurant)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*restaurant)

	db, _ := DbInterface.GetDbData()

	var id int
	var name, location string
	var isVeg bool

	query := `
	SELECT 
		restaurant_data.uid, restaurant_data.name, restaurant_data.location, restaurant_data.is_veg
	FROM 
		restaurant_data 
	WHERE 
		restaurant_data.uid = $1`

	row := db.QueryRow(query, restaurant.ID)
	err = row.Scan(&id, &name, &location, &isVeg)
	if err != nil {
		fmt.Println("Error: ", err)
		return c.Status(404).JSON(fiber.Map{
			"error": "Restaurant not found",
		})
	}

	fmt.Println(id, name, location, isVeg)

	return c.JSON(fiber.Map{
		"result": map[string]interface{}{
			"uid":      id,
			"name":     name,
			"location": location,
			"is_veg":   isVeg,
		},
		"status": "ok",
	})
}
