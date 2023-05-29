package animals

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Animal struct {
	ID        int       `json:"id" db:"id" valid:"-"`
	Id_vet    int       `json:"id_vet" db:"id_vet" valid:"-"`
	Usuario   string    `json:"usuario" db:"usuario" valid:"required"`
	Password  string    `json:"password" db:"password" valid:"required"`
	Nombres   string    `json:"nombres" db:"nombres" valid:"required"`
	Raza      string    `json:"apellidos" db:"apellidos" valid:"required"`
	Edad      string    `json:"edad" db:"edad" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewAnimal(Id_vet int, Usuario string, Password string, Nombres string, Raza string, Edad string) *Animal {
	return &Animal{
		Id_vet:   Id_vet,
		Usuario:  Usuario,
		Password: Password,
		Nombres:  Nombres,
		Raza:     Raza,
		Edad:     Edad,
	}
}

func (m *Animal) Validate() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
