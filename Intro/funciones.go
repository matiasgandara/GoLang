package main
 import "fmt"

type close func(s string)

func invokeClose(c close){
	c("Hellow GO!!!")
}

func main() {
/* 	f := func(s string){
		fmt.Println(s)
	}
	invokeClose(f) */
	//hace lo mismo... implementacion mas ordenada
	invokeClose(func(s string){
		fmt.Println(s)
	})
}