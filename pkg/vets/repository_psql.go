package vets

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/internal/models"
)

type psqldb struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func NewVetPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psqldb {
	return &psqldb{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psqldb) Create(m *Vet) error {
	const sqlInsert = `INSERT INTO public.vets (nombre,created_at, updated_at) VALUES (:nombre, Now(), Now()) `
	_, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't insert Vet: %v", err)
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psqldb) Update(m *Vet) error {
	const sqlUpdate = `UPDATE public.vets SET nombre = :nombre , updated_at = Now() WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't update Vet: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psqldb) Delete(id int) error {
	const sqlDelete = `DELETE FROM public.vets WHERE id = :id `
	m := Vet{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't delete Vet: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID obtiene un registro de la BD
func (s *psqldb) GetByID(id int) (*Vet, error) {
	const sqlGetByID = `SELECT id, nombre, created_at, updated_at FROM public.vets WHERE id = $1 `
	mdl := Vet{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't get Vet: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll obtiene todos los registros de la BD
func (s *psqldb) GetAll() ([]*Vet, error) {
	const sqlGetAll = `SELECT id, nombre, created_at, updated_at FROM public.vets `
	ms := []*Vet{}
	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't get all Vet: %v", err)
		return ms, err
	}
	return ms, nil
}
