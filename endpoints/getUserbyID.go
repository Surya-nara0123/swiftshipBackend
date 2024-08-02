package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/helperfunction"
	"github.com/surya-nara0123/swiftship/types"
	"golang.org/x/crypto/bcrypt"
)

func GetUserbyID(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
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
	helperfunction.CheckError(err)

	fmt.Println("Connected!")

	var id int
	var name, email string
	var mobile, userType int
	var hashedPassword string
	query := `
	SELECT 
		user_details.uid, user_details.username, user_details.email, user_details.mobile, 
		user_details.user_type, auth_details.password 
	FROM 
		user_details 
	INNER JOIN 
		auth_details 
	ON 
		user_details.uid = auth_details.user_id 
	WHERE 
		user_details.uid = $1`

	row := db.QueryRow(query, user.ID)
	err = row.Scan(&id, &name, &email, &mobile, &userType, &hashedPassword)
	helperfunction.CheckError(err)

	fmt.Println(id, name, email, mobile, userType)

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
