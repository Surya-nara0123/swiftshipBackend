package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

// UpdateUser is an endpoint that updates a user's information in the database
func UpdateUser(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	updateduser := new(types.UserUpdateReq)

	err := c.BodyParser(updateduser)
	if err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*updateduser)

	db, _ := DbInterface.GetDbData()

	userData := types.UserDetails{}

	db.First(&userData, "username = ?", updateduser.OriginalUsername)

	if userData.Username == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	userData.Username = updateduser.Username
	userData.Email = updateduser.Email
	userData.Mobile = updateduser.Mobile

	db.Where("username = ?", updateduser.OriginalUsername).Save(&userData)

	return c.Status(200).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}
