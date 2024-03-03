package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hasanardhian8/go-fiber-postgres/config"
	"github.com/hasanardhian8/go-fiber-postgres/models"
)

type inputPen struct {
	Nama string `json:"nama"`
	Rt   string `json:"rt"`
}

func GetPenduduk(c *fiber.Ctx) error {
	var pen []models.Penduduks

	config.DBConn.Find(&pen)
	return c.Status(fiber.StatusOK).JSON(&pen)
}

func CreatePenduduk(c *fiber.Ctx) error {
	body := inputPen{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var pen models.Penduduks
	pen.Nama = body.Nama
	pen.Rt = body.Rt

	if result := config.DBConn.Create(&pen); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(&pen)
}

func GetPendudukById(c *fiber.Ctx) error {
	id := c.Params("id")
	var pen models.Penduduks

	if result := config.DBConn.First(&pen, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&pen)
}

func UpdatePenduduk(c *fiber.Ctx) error {
	id := c.Params("id")
	updatebody := inputPen{}

	if err := c.BodyParser(&updatebody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var pen models.Penduduks

	if result := config.DBConn.First(&pen, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	pen.Nama = updatebody.Nama
	pen.Rt = updatebody.Rt

	config.DBConn.Save(&pen)

	return c.Status(fiber.StatusOK).SendString("Penduduk updated")
}

func DeletePenduduk(c *fiber.Ctx) error {
	id := c.Params("id")
	var pen []models.Penduduks

	if result := config.DBConn.First(&pen, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	config.DBConn.Delete(&pen)

	return c.Status(fiber.StatusOK).SendString("Penduduk deleted")
}
