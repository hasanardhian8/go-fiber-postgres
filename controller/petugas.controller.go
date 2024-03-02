package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hasanardhian8/go-fiber-postgres/models"
	"github.com/jinzhu/gorm"
)

func GetPetugas(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var pet []models.Petugases

	db.Find(&pet)
	return c.Status(fiber.StatusOK).JSON(&pet)
}

func CreatePetugas(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var inputPet models.Petugases

	if err := c.BodyParser(&inputPet); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	db.Create(&inputPet)
	return c.Status(fiber.StatusCreated).JSON(&inputPet)
}

func GetPetugasById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var pet models.Petugases

	if err := db.Where("id = ?", id).First(&pet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).SendString("petugas not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&pet)
}

func UpdatePetugas(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var pet models.Petugases

	if err := db.Where("id = ?", id).First(&pet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).SendString("petugas not found")
		}
		return err
	}

	if err := c.BodyParser(&pet); err != nil {
		return err
	}

	db.Save(&pet)

	return c.Status(fiber.StatusOK).SendString(" updated")
}

func DeletePetugas(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var pet []models.Petugases

	if err := db.Where("id = ?", id).First(&pet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).SendString(" not found")
		}
		return err
	}

	db.Delete(&pet)

	return c.Status(fiber.StatusOK).SendString(" deleted")
}
