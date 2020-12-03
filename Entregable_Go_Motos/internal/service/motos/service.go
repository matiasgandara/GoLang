package motos

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/matiasgandara/GoLang/tree/main/Entregable_Go_Motos/internal/config"
)

//Motos...
type Moto struct {
	ID         int
	Patente    string
	Modelo     int
	Nombre     string
	Cilindrada int
	Color      string
}

//NewMoto...
func NewMoto(ID int, Patente string, Modelo int,
	Nombre string, Cilindrada int, Color string) (Moto, error) {
	return Moto{ID, Patente, Modelo, Nombre, Cilindrada, Color}, nil
}

// Service...
type Service interface {
	AddMoto(Moto) error
	FindByID(int) (*Moto, error)
	FindAll() ([]*Moto, error)
	Delete(int) error
	Update(Moto) error
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// AddMoto ...
func (s service) AddMoto(m Moto) error {
	query := `INSERT INTO motos (
		patente, modelo, nombre, cilindrada, color) VALUES (?, ?, ?, ?, ?)`

	_, err := s.db.Exec(query, m.Patente, m.Modelo, m.Nombre, m.Cilindrada, m.Color)

	if err != nil {
		return err
	}
	return nil
}

// FindByID ...
func (s service) FindByID(ID int) (*Moto, error) {
	var f []*Moto
	if err := s.db.Select(&f, "SELECT * FROM motos WHERE id = ?", ID); err != nil {
		return nil, err
	}

	if len(f) > 0 {
		return f[0], nil
	} else {
		return nil, nil
	}
}

// FindAll ...
func (s service) FindAll() ([]*Moto, error) {
	var list []*Moto
	if err := s.db.Select(&list, "SELECT * FROM motos"); err != nil {
		return nil, err
	}
	return list, nil
}

// Delete ...
func (s service) Delete(ID int) error {
	query := "DELETE FROM motos WHERE id = ?"

	_, err := s.db.Exec(query, ID)

	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (s service) Update(m Moto) error {
	query := `UPDATE motos 
		SET patente = ?, modelo = ?, nombre = ?, cilindrada = ?, color = ? 
		WHERE id = ?`

	fmt.Println(m.Patente)

	_, err := s.db.Exec(query, m.Patente,
		m.Modelo, m.Nombre, m.Cilindrada, m.Color, m.ID)

	if err != nil {
		return err
	}
	return nil
}
