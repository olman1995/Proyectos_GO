package main

import(
	"fmt" //imprimir
	"strconv" //conversiones
) 

func main(){
	//las variables inician con un valor predeterminado
	//var x,y,z int // int = 0
	//var c string // string = ""
	//var b bool // bool = false
	//x := 23 tipado dinamico
	//nombre := "abc"
	edad := 22
	edad_str := strconv.Itoa(edad)
	fmt.Println("Mi edad es "+edad_str)
	x:=30
	y_str:="30"
	y,_:=strconv.Atoi(y_str) // _ ignorar variables
	fmt.Println(x+y)
}