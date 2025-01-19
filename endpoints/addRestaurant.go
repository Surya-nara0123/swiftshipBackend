package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/helperfunction"
	"github.com/surya-nara0123/swiftship/types"

	"golang.org/x/crypto/bcrypt"
)

func AddRestaurant(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	restaurant := new(types.RestuarantReq)

	err := c.BodyParser(restaurant)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*restaurant)

	//generate unique id
	uid := helperfunction.GenerateUniqueInt()
	fmt.Println(uid)

	db, _ := DbInterface.GetDbData()

	restaurantData := &types.RestaurantData{
		UID:        uid,
		Name:       restaurant.Name,
		Location:   restaurant.Location,
		IsVeg:      restaurant.IsVeg,
		VendorName: restaurant.VendorName,
		StatusId:   false,
	}

	res := db.Create(restaurantData)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(restaurant.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// insert the auth details in the auth_details table
	authDetails := &types.AuthDetails{
		UserID:   uid,
		Password: string(hashedPassword),
	}

	result := db.Create(authDetails)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	fmt.Println("Restaurant created!")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
