package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hasanardhian8/go-fiber-postgres/config"
	"github.com/hasanardhian8/go-fiber-postgres/routes"
)

func main() {
	app := fiber.New()
	config.DatabaseConnection()

	fmt.Println("Connected to database")
	routes.RouterInit(app)
	app.Listen(":8080")

	//defer config.DBConn.Close()

}
