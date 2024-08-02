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
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*user)

	db, _ := DbInterface.GetDbData()

	// generate unique id
	uid := helperfunction.GenerateUniqueInt()

	err = db.Ping()
	helperfunction.CheckError(err)

	fmt.Println("Connected!")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	helperfunction.CheckError(err)
	fmt.Println(string(hashedPassword))

	_, err = db.Exec("INSERT INTO user_details (uid, username, email, mobile, user_type) VALUES ($1, $2, $3, $4, $5)", uid, user.Name, user.Email, user.Mobile, user.UserType)
	helperfunction.CheckError(err)

	_, err = db.Exec("INSERT INTO auth_details (user_id, password) VALUES ($1, $2)", uid, string(hashedPassword))
	helperfunction.CheckError(err)

	fmt.Println("User created!")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
