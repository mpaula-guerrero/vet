package sessions_handler

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/pkg/animals"
)

type Handler struct {
	DB   *sqlx.DB
	TxID string
}

func (h *Handler) login(c *fiber.Ctx) error {
	res := LoginResponse{}
	m := LoginRequest{}

	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo leer el Modelo crear usuario: %v", err)
		return c.Status(http.StatusForbidden).JSON(res)
	}

	serviceAnimal := animals.NewAnimalService(animals.NewAnimalPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	animal, err := serviceAnimal.GetByAnimal(m.Usuario)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusForbidden).JSON(res)
		}
		logger.Error.Printf(h.TxID, "no se pudo obtener el animale: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if m.Password != animal.Password {
		return c.Status(http.StatusForbidden).JSON(res)
	}
	animal.Password = ""
	return c.Status(http.StatusOK).JSON(animal)

}
