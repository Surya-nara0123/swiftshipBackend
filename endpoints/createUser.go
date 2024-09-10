package endpoints

import (
	"fmt"

	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/helperfunction"
	"github.com/surya-nara0123/swiftship/types"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	user := new(types.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(*user)

	db, _ := DbInterface.GetDbData()

	//check if the user is a valid request
	if user.Name == "" || user.Email == "" || user.Mobile == 0 || user.UserType == 0 || user.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	// generate unique id
	uid := helperfunction.GenerateUniqueInt()

	fmt.Println("Connected!")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println(string(hashedPassword))

	// create a new user record
	newUser := &types.UserDetails{
		UID:      uid,
		Username: user.Name,
		Email:    user.Email,
		Mobile:   user.Mobile,
		UserType: user.UserType,
	}

	if user.UserType == 0 {
		user.UserType = 1
	}
	// insert the new user record into the user_details table
	result := db.Create(newUser)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return c.Status(400).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// insert the auth details in the auth_details table
	authDetails := &types.AuthDetails{
		UserID:   uid,
		Password: string(hashedPassword),
	}

	result = db.Create(authDetails)

	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	fmt.Println("User created!")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
