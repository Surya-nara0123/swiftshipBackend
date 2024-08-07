package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetUserID(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	user := new(types.UserGet)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*user)

	db, _ := DbInterface.GetDbData()

	err = db.Ping()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println("Connected!")

	var id int
	query := `
	SELECT 
		uid
	FROM 
		user_details 
	WHERE 
		username = $1`

	row := db.QueryRow(query, user.ID)
	err = row.Scan(&id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println(id)
	return c.JSON(fiber.Map{
		"uid": id,
	})
}
