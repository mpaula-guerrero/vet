package animales_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func AnimalesRouter(app *fiber.App, db *sqlx.DB, tx string) {

	animalHd := Handler{DB: db, TxID: tx}

	api := app.Group("/api")
	v1 := api.Group("/v1/animales")
	v1.Post("/", animalHd.create)
	v1.Put("/:id", animalHd.update)
	v1.Delete("/:id", animalHd.delete)
	v1.Get("/:id", animalHd.getByID)
	v1.Get("/", animalHd.getAll)

}
