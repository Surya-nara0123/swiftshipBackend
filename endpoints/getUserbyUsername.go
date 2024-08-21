package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
	"golang.org/x/crypto/bcrypt"
)

func GetUserbyUsername(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	user := new(types.UserUsernameReq)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*user)

	db, _ := DbInterface.GetDbData()

	var hashedPassword string
	userDetails := new(types.UserDetails)

	a := db.First(&userDetails, "username = ?", user.Username)
	authDetails := new(types.AuthDetails)
	restaurant := types.RestaurantData{}
	fmt.Println("hiiii", *a)
	if a.Error != nil {
		fmt.Printf("User not found! %v\n", a.Error)
		db.First(&restaurant, "vendor_name = ?", user.Username)
		fmt.Println("restaurant", restaurant)
		db.First(&authDetails, restaurant.UID)
	} else {
		db.First(&authDetails, userDetails.UID)
	}

	hashedPassword = authDetails.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		fmt.Printf("Password incorrect! %v\n", err)
		return c.Status(401).JSON(fiber.Map{
			"error": "Password incorrect",
		})
	}

	fmt.Println("Password correct!")
	if a.Error == nil {
		return c.JSON(fiber.Map{
			"status": "ok",
			"user": fiber.Map{
				"id":        userDetails.UID,
				"name":      userDetails.Username,
				"email":     userDetails.Email,
				"mobile":    userDetails.Mobile,
				"user_type": userDetails.UserType,
			},
		})
	} else {
		return c.JSON(fiber.Map{
			"status": "ok",
			"resturant": fiber.Map{
				"id":          restaurant.UID,
				"name":        restaurant.Name,
				"location":    restaurant.Location,
				"is_veg":      restaurant.IsVeg,
				"vendor_name": restaurant.VendorName,
				"user_type":   2,
			},
		})
	}
}
