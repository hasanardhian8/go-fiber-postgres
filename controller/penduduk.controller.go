package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hasanardhian8/go-fiber-postgres/models"
	"github.com/jinzhu/gorm"
)

func GetPenduduk(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var pen []models.Penduduks

	db.Find(&pen)
	return c.Status(fiber.StatusOK).JSON(&pen)
}

func CreatePenduduk(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var inputPen models.Penduduks

	if err := c.BodyParser(&inputPen); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	db.Create(&inputPen)
	return c.Status(fiber.StatusCreated).JSON(&inputPen)
}

func GetPendudukById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var pen models.Penduduks

	if err := db.Where("id = ?", id).First(&pen).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).SendString("Penduduk not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&pen)
}

func UpdatePenduduk(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var pen models.Penduduks

	if err := db.Where("id = ?", id).First(&pen).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).SendString("Penduduk not found")
		}
		return err
	}

	if err := c.BodyParser(&pen); err != nil {
		return err
	}

	db.Save(&pen)

	return c.Status(fiber.StatusOK).SendString("Penduduk updated")
}

func DeletePenduduk(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var pen []models.Penduduks

	if err := db.Where("id = ?", id).First(&pen).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).SendString("Penduduk not found")
		}
		return err
	}

	db.Delete(&pen)

	return c.Status(fiber.StatusOK).SendString("Penduduk deleted")
}
