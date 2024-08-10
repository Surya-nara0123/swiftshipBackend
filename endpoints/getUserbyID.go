package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
	"golang.org/x/crypto/bcrypt"
)

func GetUserbyID(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	user := new(types.UserIdReq)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*user)

	db, _ := DbInterface.GetDbData()

	fmt.Println("Connected!")

	var id int64
	var name, email string
	var mobile int64
	var userType int
	var hashedPassword string

	authDetails := new(types.AuthDetails)

	db.First(&authDetails, user.ID)

	userDetails := new(types.UserDetails)

	db.First(&userDetails, user.ID)

	fmt.Println(userDetails)

	fmt.Println(id, name, email, mobile, userType)

	id = userDetails.UID
	name = userDetails.Username
	email = userDetails.Email
	mobile = userDetails.Mobile
	userType = userDetails.UserType
	hashedPassword = authDetails.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		fmt.Printf("Password incorrect! %v\n", err)
		return c.Status(401).JSON(fiber.Map{
			"error": "Password incorrect",
		})
	}

	fmt.Println("Password correct!")
	return c.JSON(fiber.Map{
		"status": "ok",
		"user": fiber.Map{
			"id":        id,
			"name":      name,
			"email":     email,
			"mobile":    mobile,
			"user_type": userType,
		},
	})
}
