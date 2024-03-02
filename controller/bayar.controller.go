package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hasanardhian8/go-fiber-postgres/config"
	"github.com/hasanardhian8/go-fiber-postgres/models"
)

type inputBay struct {
	Id         int       `json:"id" `
	Idpenduduk int       `json:"idPenduduk" `
	Idpetugas  int       `json:"idPetugas" `
	Nominal    int       `json:"nominal"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func GetBayar(c *fiber.Ctx) error {
	var bay []models.Bayars

	config.DBConn.Find(&bay)
	return c.Status(fiber.StatusOK).JSON(&bay)
}

func CreateBayar(c *fiber.Ctx) error {
	body := inputBay{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var bay models.Bayars
	bay.Idpenduduk = body.Idpenduduk
	bay.Idpetugas = body.Idpetugas
	bay.Nominal = body.Nominal

	if result := config.DBConn.Create(&bay); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&bay)
}

func GetBayarById(c *fiber.Ctx) error {
	id := c.Params("Id")

	var bay models.Bayars

	if result := config.DBConn.First(&bay, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&bay)
}

func UpdateBayar(c *fiber.Ctx) error {
	id := c.Params("Id")
	updatebody := inputBay{}

	if err := c.BodyParser(&updatebody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var bay models.Bayars

	if result := config.DBConn.First(&bay, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	bay.Idpenduduk = updatebody.Idpenduduk
	bay.Idpetugas = updatebody.Idpetugas
	bay.Nominal = updatebody.Nominal

	config.DBConn.Save(&bay)

	return c.Status(fiber.StatusOK).SendString("Bayar updated")
}

func DeleteBayar(c *fiber.Ctx) error {
	id := c.Params("Id")

	var bay []models.Bayars

	if result := config.DBConn.First(&bay, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	config.DBConn.Delete(&bay)

	return c.Status(fiber.StatusOK).SendString("Bayar deleted")
}
