package main

import "fmt"

//Printable....
type Printable interface{
	print()
}
// modelo de moto
type Moto struct {
	marca string
	patente string
	cilindrada int
	modelo int
}
// modeloado de agencia
type Agencia struct{
	motos []Moto
	nombre string
	localidad string
}
// invoca Printeable
func invokePrint(p Printable) {
	p.print()
}

func newMoto(marca, patente string, modelo, cilindrada int) Moto{
	return Moto{
		marca: marca,
		patente: patente,
		cilindrada: cilindrada,
		modelo: modelo,
	}
}

func newAgencia(nombre, localidad string, motos []Moto) Agencia{
	return Agencia{
		motos: motos,
		nombre: nombre, 
		localidad: localidad,
	}
}

func (a *Agencia) print() {
	fmt.Println("[Agencia] nombre", a.nombre) 
	fmt.Println("localidad",a.localidad) 
	fmt.Println("Motos", a.motos)
}

func (m *Moto) print() {
	fmt.Println("[Moto] marca", m.marca) 
	fmt.Println("patente",m.patente) 
	fmt.Println("cilindrada",m.cilindrada)
	fmt.Println("modelo", m.modelo)
}


func main(){
	m := newMoto("Honda","123asd",2500,2015)
	motos := make([]Moto,5)
	motos= append(motos,m)
	a := newAgencia("Roads","Tandil", motos)
	m.print()
	a.print()
	/* invokePrint(m) */
}