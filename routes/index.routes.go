package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hasanardhian8/go-fiber-postgres/controller"
)

func RouterInit(r *fiber.App) {

	r.Get("/penduduk", controller.GetPenduduk)
	r.Post("/penduduk", controller.CreatePenduduk)
	r.Get("/penduduk/:id", controller.GetPendudukById)
	r.Patch("/penduduk/:id", controller.UpdatePenduduk)
	r.Delete("/penduduk/:id", controller.DeletePenduduk)

	r.Get("/bayar", controller.GetBayar)
	r.Post("/bayar", controller.CreateBayar)
	r.Get("/bayar/:id", controller.GetBayarById)
	r.Patch("/bayar/:id", controller.UpdateBayar)
	r.Delete("/bayar/:id", controller.DeleteBayar)

	r.Get("/petugas", controller.GetPetugas)
	r.Post("/petugas", controller.CreatePetugas)
	r.Get("/petugas/:id", controller.GetPetugasById)
	r.Patch("/petugas/:id", controller.UpdatePetugas)
	r.Delete("/petugas/:id", controller.DeletePetugas)

}
