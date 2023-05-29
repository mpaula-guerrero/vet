package vets

import (
	"github.com/jmoiron/sqlx"
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesVetRepository interface {
	Create(m *Vet) error
	Update(m *Vet) error
	Delete(id int) error
	GetByID(id int) (*Vet, error)
	GetAll() ([]*Vet, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesVetRepository {
	var s ServicesVetRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return NewVetPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
