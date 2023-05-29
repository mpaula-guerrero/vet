package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
	"test_ecapture_backend/api/rest/handler/animales_handler"
	"test_ecapture_backend/api/rest/handler/sessions_handler"
	"test_ecapture_backend/api/rest/handler/vets_handler"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/jmoiron/sqlx"
)

func routes(db *sqlx.DB, loggerHttp bool, allowedOrigins string) *fiber.App {
	app := fiber.New()

	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST",
	}))
	if loggerHttp {
		app.Use(logger.New())
	}
	TxID := uuid.New().String()

	vets_handler.VetsRouter(app, db, TxID)
	sessions_handler.SessionRouter(app, db, TxID)
	animales_handler.AnimalesRouter(app, db, TxID)
	return app
}
