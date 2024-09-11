package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/helperfunction"
	"github.com/surya-nara0123/swiftship/types"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// CreateNormalUser is a function that creates a new user in the database
func CreateNormalUser(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	// Get the user details from the request body
	user := new(types.NormalUser)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Could not parse request body",
		})
	}

	db, _ := DbInterface.GetDbData()

	if user.Name == "" || user.Email == "" || user.Mobile == 0 || user.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	uid := helperfunction.GenerateUniqueInt()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Could not hash password",
		})
	}

	newUser := &types.UserDetails{
		UID:      uid,
		Username: user.Name,
		Email:    user.Email,
		Mobile:   user.Mobile,
		UserType: 1,
	}

	result := db.Create(newUser)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return c.Status(400).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

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
