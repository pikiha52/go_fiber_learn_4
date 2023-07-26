package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"

	"relationship_go/database"
)

type User struct {
	ID          string
	Name        string
	Username    string
	Password    string
	UserAddress []Address
	Role        []Role `gorm:"many2many:user_roles;"`
}

type Address struct {
	ID      string
	UserID  string
	Address string
}

type Addresses struct {
	ID      string
	UserID  string
	Address string
	User    User
}

type Role struct {
	ID   string
	Name string
}

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/hasOne", func(c *fiber.Ctx) error {
		db := database.DB
		var address []Addresses

		db.Model(&address).Preload(clause.Associations).Find(&address)

		return c.JSON(fiber.Map{"status": "oke", "data": address})
	})

	app.Get("/hasMany", func(c *fiber.Ctx) error {
		db := database.DB
		var users []User

		db.Model(&users).Preload(clause.Associations).Find(&users)

		return c.JSON(fiber.Map{"status": "oke", "data": users})
	})

	app.Get("/manyToMany", func(c *fiber.Ctx) error {
		db := database.DB
		var users []User
		db.Preload(clause.Associations).Find(&users)

		return c.JSON(fiber.Map{"status": "oke", "data": users})
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3000")
}
