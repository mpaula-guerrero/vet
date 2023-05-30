package vets

import (
	"vet_ecapture_backend/internal/logger"
	"vet_ecapture_backend/internal/models"
)

type ServicesVet interface {
	Create(nombre string) error
	Update(id int, nombre string) error
	Delete(id int) error
	GetByID(id int) (*Vet, error)
	GetAll() ([]*Vet, error)
}

type service struct {
	repository ServicesVetRepository
	user       *models.User
	txID       string
}

func (s service) Create(nombre string) error {

	m := NewVet(nombre)

	valid, err := m.Validate()
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return err
	}

	if err := s.repository.Create(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't create Vet :", err)
		return err
	}
	return nil
}

func (s service) Update(id int, nombre string) error {
	m := NewVet(nombre)
	m.ID = id
	valid, err := m.Validate()
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return err
	}

	if err := s.repository.Update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Vet :", err)
		return err
	}
	return nil
}

func (s service) Delete(id int) error {
	v, err := s.repository.GetByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Vet :", err)
		return err
	}
	if v == nil {
		logger.Error.Println(s.txID, " - couldn't get Vet %d to delete:", id)
		return err
	}
	if err := s.repository.Delete(id); err != nil {
		logger.Error.Println(s.txID, " - couldn't delete Vet :", err)
		return err
	}
	return nil
}

func (s service) GetByID(id int) (*Vet, error) {
	v, err := s.repository.GetByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Vet :", err)
		return nil, err
	}
	return v, nil
}

func (s service) GetAll() ([]*Vet, error) {
	v, err := s.repository.GetAll()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Vets :", err)
		return nil, err
	}
	return v, nil
}

func NewVetService(repository ServicesVetRepository, user *models.User, txID string) ServicesVet {
	return &service{
		repository: repository,
		user:       user,
		txID:       txID,
	}
}
