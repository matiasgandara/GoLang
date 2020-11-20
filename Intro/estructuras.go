package main

import "fmt"

//Printable....
type Printable interface{
	print()
}

type person struct {
	name string
}

type figure struct {
	name string
}

func (f *figure) print() {
	fmt.Println("[figure]", f.name)
}

func (p *person) print() {
	fmt.Println("[person]", p.name)
}

func invokePrint(p Printable) {
	p.print()
}


func main(){
	p := &person{name: "Juan"}
	invokePrint(p)

	f := &figure{name: "Triangulo"}
	invokePrint(f)
}