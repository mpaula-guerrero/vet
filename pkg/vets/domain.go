package vets

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Vet struct {
	ID        int       `json:"id" db:"id" valid:"-"`
	Nombre    string    `json:"nombre" db:"nombre" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt  time.Time `json:"update_at" db:"update_at"`
}

func NewVet(nombre string) *Vet {
	return &Vet{
		Nombre: nombre,
	}
}

func (m *Vet) Validate() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
