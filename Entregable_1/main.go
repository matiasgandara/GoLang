package main

import (
	"fmt"
)

func main() {

	//CRUD...
	c := NewConcesionario("Todo_Motos", "Tandil")
	c.Add(Moto{0, "Yamaha", "123ABC", 1500, 2015})
	c.Add(Moto{1, "Honda", "888ABC", 250, 2017})
	c.Add(Moto{2, "Ducati", "456ABC", 2000, 2010})
	c.Add(Moto{3, "Ciambretta", "abcd1234", 100, 1987})

	c.Print()

	m0 := c.FindByID(0)
	if m0 != nil {
		fmt.Println("Se encontro el ID=0")
	}

	c.Delete(0)
	c.Print()

	c.Update(Moto{1, "Hondaaaa", "888ABC", 200, 2017})
	c.Print()
}
