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
	AddMoto(Moto)
	FindByID(int) *Moto
	FindAll() []*Moto
	Delete(int)
	Update(Moto)
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
func (s service) AddMoto(m Moto) {
	query := `INSERT INTO motos (
		patente, modelo, nombre, cilindrada, color) VALUES (?, ?, ?, ?, ?)`

	_, err := s.db.Exec(query, m.Patente, m.Modelo, m.Nombre, m.Cilindrada, m.Color)

	if err != nil {
		panic(err)
	}
}

// FindByID ...
func (s service) FindByID(ID int) *Moto {
	var f []*Moto
	if err := s.db.Select(&f, "SELECT * FROM motos WHERE id = ?", ID); err != nil {
		panic(err)
	}

	if len(f) > 0 {
		return f[0]
	} else {
		return nil
	}
}

// FindAll ...
func (s service) FindAll() []*Moto {
	var list []*Moto
	if err := s.db.Select(&list, "SELECT * FROM motos"); err != nil {
		panic(err)
	}
	return list
}

// Delete ...
func (s service) Delete(ID int) {
	query := "DELETE FROM motos WHERE id = ?"

	_, err := s.db.Exec(query, ID)

	if err != nil {
		panic(err)
	}
}

// Update ...
func (s service) Update(m Moto) {
	query := `UPDATE motos 
		SET patente = ?, modelo = ?, nombre = ?, cilindrada = ?, color = ? 
		WHERE id = ?`

	fmt.Println(m.Patente)

	_, err := s.db.Exec(query, m.Patente,
		m.Modelo, m.Nombre, m.Cilindrada, m.Color, m.ID)

	if err != nil {
		panic(err)
	}
}
