package animales_handler

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/pkg/animals"
)

type Handler struct {
	DB   *sqlx.DB
	TxID string
}

func (h *Handler) create(c *fiber.Ctx) error {
	res := AnimalResponse{}
	m := AnimalRequest{}

	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo leer el Modelo crear animal: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	serviceAnimal := animals.NewAnimalService(animals.NewAnimalPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)

	err = serviceAnimal.Create(m.Id_vet, m.Usuario, m.Password, m.Nombres, m.Raza, m.Edad)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo crear el animal: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	return c.Status(http.StatusCreated).JSON(AnimalResponse{
		Id_vet:   m.Id_vet,
		Usuario:  m.Usuario,
		Password: m.Password,
		Nombres:  m.Nombres,
		Raza:     m.Raza,
		Edad:     m.Edad,
	})
}

func (h *Handler) update(c *fiber.Ctx) error {
	res := AnimalResponse{}
	m := AnimalRequest{}
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	err = c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo leer el Modelo crear animal: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	fmt.Println(m)
	serviceAnimal := animals.NewAnimalService(animals.NewAnimalPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	err = serviceAnimal.Update(id, m.Id_vet, m.Usuario, m.Password, m.Nombres, m.Raza, m.Edad)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo actualizar el animal: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	return c.Status(http.StatusOK).JSON(AnimalResponse{
		Id_vet:   m.Id_vet,
		Usuario:  m.Usuario,
		Password: m.Password,
		Nombres:  m.Nombres,
		Raza:     m.Raza,
		Edad:     m.Edad,
	})

}

func (h *Handler) delete(c *fiber.Ctx) error {
	res := AnimalResponse{}
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	serviceAnimal := animals.NewAnimalService(animals.NewAnimalPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	err = serviceAnimal.Delete(id)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo eliminar el Animal: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(res)

}

func (h *Handler) getByID(c *fiber.Ctx) error {
	res := AnimalResponse{}

	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	serviceAnimal := animals.NewAnimalService(animals.NewAnimalPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	p, err := serviceAnimal.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusOK).JSON(res)
		}
		logger.Error.Printf(h.TxID, "no se pudo obtener el Animal: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(p)

}

func (h *Handler) getAll(c *fiber.Ctx) error {
	var res []AnimalResponse

	serviceAnimal := animals.NewAnimalService(animals.NewAnimalPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	ps, err := serviceAnimal.GetAll()
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo obtener los Animals: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(ps)

}
