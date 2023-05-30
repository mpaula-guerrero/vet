package vets_handler

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
	"vet_ecapture_backend/internal/logger"
	"vet_ecapture_backend/pkg/vets"
)

type Handler struct {
	DB   *sqlx.DB
	TxID string
}

func (h *Handler) create(c *fiber.Ctx) error {
	res := VetResponse{}
	m := VetRequest{}

	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo leer el Modelo crear vet: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	serviceVet := vets.NewVetService(vets.NewVetPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)

	err = serviceVet.Create(m.Nombre)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo crear el vet: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	res.Nombre = m.Nombre
	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) update(c *fiber.Ctx) error {
	res := VetResponse{}
	m := VetRequest{}
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	err = c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo leer el Modelo crear Vet: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	serviceVet := vets.NewVetService(vets.NewVetPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	err = serviceVet.Update(id, m.Nombre)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo actualizar el vet: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	res.Nombre = m.Nombre
	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) delete(c *fiber.Ctx) error {
	res := VetResponse{}
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	serviceVet := vets.NewVetService(vets.NewVetPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	err = serviceVet.Delete(id)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo eliminar el vet: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(res)

}

func (h *Handler) getByID(c *fiber.Ctx) error {
	res := VetResponse{}

	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	serviceVet := vets.NewVetService(vets.NewVetPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	p, err := serviceVet.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusOK).JSON(res)
		}
		logger.Error.Printf(h.TxID, "no se pudo obtener el Vet: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(p)

}

func (h *Handler) getAll(c *fiber.Ctx) error {
	var res []VetResponse

	serviceVet := vets.NewVetService(vets.NewVetPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	ps, err := serviceVet.GetAll()
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo obtener los Vets: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(ps)

}
