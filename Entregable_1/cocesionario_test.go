package main

import "testing"

func TestConcesionarioAdd(t *testing.T) {
	c := NewConcesionario("Todo_Motos", "Rauch")
	m0 := c.FindByID(0)
	if m0 != nil {
		t.Error("La moto con ID=0 ya existe")
	}

	c.Add(Moto{0, "Yamaha", "123ABC", 1500, 2015})
	m0 = c.FindByID(0)
	if m0 == nil {
		t.Error("La moto con ID = 0 no fue agregada")
	}

	if m0.Marca != "Yamaha" {
		t.Error("La moto con ID =0  no tiene la marca correcta")
	}
}
