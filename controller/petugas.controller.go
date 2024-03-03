package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hasanardhian8/go-fiber-postgres/config"
	"github.com/hasanardhian8/go-fiber-postgres/models"
)

type inputPet struct {
	Nama     string `json:"nama"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetPetugas(c *fiber.Ctx) error {
	var pet []models.Petugases

	config.DBConn.Find(&pet)
	return c.Status(fiber.StatusOK).JSON(&pet)
}

func CreatePetugas(c *fiber.Ctx) error {
	body := inputPet{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var pet models.Petugases
	pet.Nama = body.Nama
	pet.Password = body.Password
	pet.Role = body.Role

	if result := config.DBConn.Create(&pet); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&pet)
}

func GetPetugasById(c *fiber.Ctx) error {
	id := c.Params("id")
	var pet models.Petugases

	if result := config.DBConn.First(&pet, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&pet)
}

func UpdatePetugas(c *fiber.Ctx) error {
	id := c.Params("id")
	updatebody := inputPet{}

	if err := c.BodyParser(&updatebody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var pet models.Petugases

	if result := config.DBConn.First(&pet, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	pet.Nama = updatebody.Nama
	pet.Password = updatebody.Password
	pet.Role = updatebody.Role

	config.DBConn.Save(&pet)

	return c.Status(fiber.StatusOK).SendString(" updated")
}

func DeletePetugas(c *fiber.Ctx) error {
	id := c.Params("id")
	var pet []models.Petugases

	if result := config.DBConn.First(&pet, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	config.DBConn.Delete(&pet)

	return c.Status(fiber.StatusOK).SendString(" deleted")
}
