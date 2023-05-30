package animals

import (
	"vet_ecapture_backend/internal/logger"
	"vet_ecapture_backend/internal/models"
)

type ServicesAnimal interface {
	Create(Id_vet int, Usuario string, Password string, Nombre string, Raza string, Edad string) error
	Update(id int, Id_vet int, Usuario string, Password string, Nombre string, Raza string, Edad string) error
	Delete(id int) error
	GetByID(id int) (*Animal, error)
	GetByAnimal(usuario string) (*Animal, error)
	GetAll() ([]*Animal, error)
}

type service struct {
	repository ServicesAnimalRepository
	user       *models.User
	txID       string
}

func (s service) Create(Id_vet int, Usuario string, Password string, Nombre string, Raza string, Edad string) error {
	m := NewAnimal(Id_vet, Usuario, Password, Nombre, Raza, Edad)
	valid, err := m.Validate()
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return err
	}

	if err := s.repository.Create(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't create Animal :", err)
		return err
	}
	return nil

}

func (s service) Update(id int, Id_vet int, Usuario string, Password string, Nombre string, Raza string, Edad string) error {
	m := NewAnimal(Id_vet, Usuario, Password, Nombre, Raza, Edad)
	m.ID = id
	valid, err := m.Validate()
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return err
	}

	if err := s.repository.Update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Animal :", err)
		return err
	}
	return nil
}

func (s service) Delete(id int) error {
	a, err := s.repository.GetByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Animal :", err)
		return err
	}
	if a == nil {
		logger.Error.Println(s.txID, " - couldn't get Animal %d to delete:", id)
		return err
	}
	if err := s.repository.Delete(id); err != nil {
		logger.Error.Println(s.txID, " - couldn't delete Animal :", err)
		return err
	}
	return nil
}

func (s service) GetByID(id int) (*Animal, error) {
	a, err := s.repository.GetByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Animal :", err)
		return nil, err
	}
	return a, nil
}

func (s service) GetByAnimal(usuario string) (*Animal, error) {
	a, err := s.repository.GetByAnimal(usuario)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Animal :", err)
		return nil, err
	}
	return a, nil
}

func (s service) GetAll() ([]*Animal, error) {
	a, err := s.repository.GetAll()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Animal :", err)
		return nil, err
	}
	return a, nil
}

func NewAnimalService(repository ServicesAnimalRepository, user *models.User, txID string) ServicesAnimal {
	return &service{
		repository: repository,
		user:       user,
		txID:       txID,
	}
}
