package main
import "fmt"
func main() {
	fmt.Printf("hello, world\n")

	var arr [2]int
	arr[0] = 1
	arr[1] = 2

	for i, v := range arr{
		fmt.Println(i, v)
	}
	

	fmt.Println("--------slides----")

	l := make([]int, 10)
	l = append(l, 10)
	fmt.Printf("%p\n",l)
	l = append(l,100)
	fmt.Printf("%p\n",l)//imprimo la direccion de memoria
	for i, v := range l{

		fmt.Println(i,v)
	}

	fmt.Println("-------Map------------")

	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"

	for k, v := range m {
		fmt.Println(k , v)
	}
}