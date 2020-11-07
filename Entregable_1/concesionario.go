package main

import (
	"fmt"
)

// Consecionario struct....
type Concesionario struct {
	motos     map[int]*Moto
	Nombre    string
	Localidad string
}

//Moto struct...
type Moto struct {
	ID         int
	Marca      string
	Patente    string
	Cilindrada int
	Modelo     int
}

//New Concesionario
func NewConcesionario(nombre, localidad string) Concesionario {
	motos := make(map[int]*Moto)
	return Concesionario{
		motos,
		nombre,
		localidad,
	}
}

//Add...
func (c Concesionario) Add(m Moto) {
	c.motos[m.ID] = &m
}

//Print...
func (c Concesionario) Print() {
	fmt.Println(c.Nombre, c.Localidad)
	for _, v := range c.motos {
		fmt.Printf("[%v]\t%v\n", v.ID, v.Marca)
	}
}

//FindID...
func (c Concesionario) FindByID(ID int) *Moto {
	return c.motos[ID]
}

//Delete moto de concesionario ...
func (c Concesionario) Delete(ID int) {
	delete(c.motos, ID)
}

//Update moto de concesionario ...
func (c Concesionario) Update(m Moto) {
	c.motos[m.ID] = &m
}
