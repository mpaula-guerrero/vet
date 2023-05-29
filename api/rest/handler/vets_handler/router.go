package vets_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func VetsRouter(app *fiber.App, db *sqlx.DB, tx string) {

	vetHd := Handler{DB: db, TxID: tx}

	api := app.Group("/api")
	v1 := api.Group("/v1/vets")
	v1.Post("/", vetHd.create)
	v1.Put("/:id", vetHd.update)
	v1.Delete("/:id", vetHd.delete)
	v1.Get("/:id", vetHd.getByID)
	v1.Get("/", vetHd.getAll)

}
