package animals

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"vet_ecapture_backend/internal/logger"
	"vet_ecapture_backend/internal/models"
)

type psqldb struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func NewAnimalPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psqldb {
	return &psqldb{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s psqldb) Create(m *Animal) error {
	const sqlInsert = `INSERT INTO public.animals (id_vet, usuario, password, nombre, raza, edad, created_at, update_at) 
						VALUES (:id_vet, :usuario, :password, :nombre, :raza, :edad, Now(), Now())`
	_, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't insert Animal: %v", err)
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s psqldb) Update(m *Animal) error {
	const sqlUpdate = `UPDATE public.animals SET id_vet = :id_vet, usuario = :usuario, password = :password, nombre = :nombre, raza = :raza, edad = :edad, update_at = Now() WHERE id = :id`
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't update Animal: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s psqldb) Delete(id int) error {
	const sqlDelete = `DELETE FROM public.animals WHERE id = :id`
	m := Animal{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't delete Animal: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID obtiene un registro de la BD
func (s psqldb) GetByID(id int) (*Animal, error) {
	const slqGetByID = `SELECT id, id_vet, usuario, password, nombre, raza, edad, created_at, update_at FROM public.animals WHERE id = $1`
	mdl := Animal{}
	err := s.DB.Get(&mdl, slqGetByID, id)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't get Animal: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

// GetByAnimal obtiene un registro de la BD
func (s *psqldb) GetByAnimal(usuario string) (*Animal, error) {
	const sqlGetByAnimal = `SELECT id, id_vet, usuario, password, nombre, raza, edad, created_at, update_at FROM public.animals WHERE usuario = $1 `
	mdl := Animal{}
	err := s.DB.Get(&mdl, sqlGetByAnimal, usuario)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't get Animal: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll obtiene todos los registros de la BD
func (s *psqldb) GetAll() ([]*Animal, error) {
	const sqlGetAll = `SELECT id, id_vet, usuario, password, nombre, raza, edad, created_at, update_at FROM public.animals `
	ms := []*Animal{}
	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't get all Animal: %v", err)
		return ms, err
	}
	return ms, nil
}
