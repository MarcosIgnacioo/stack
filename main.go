package main

import (
	"fmt"

	"github.com/MarcosIgnacioo/stack"
)

func main() {
	var input int
	s := stack.NewStack()
	for {
		fmt.Println(`Elija la opcion que quiera hacer
1.- Llenar el stack (100 elementos)
2.- Vaciar el stack (100 elementos)
3.- Invertir el stack (100 elementos)
4.- Sacar el elemento en el tope del stack
5.- Imprimir el stack
6.- Checar si esta vacio
7.- Checar si esta lleno
0.- Salir
		`)
		fmt.Scanln(&input)
		if input == 0 {
			break
		}
		switch input {
		case 1:
			fmt.Println("Llenando el stack")
			s.FillStack()
			fmt.Println("Stack llenado")
		case 2:
			fmt.Println("Vaciando el stack")
			s.Empty()
			fmt.Println("Stack vaciado")
		case 3:
			fmt.Println("Invirtiendo el stack")
			s.Reverse()
			fmt.Println("Stack invertido")
		case 4:
			fmt.Println("Pop: ", s.Pop())
		case 5:
			s.Print()
		case 6:
			fmt.Println("El stack esta vacio?", s.IsEmpty())
		case 7:
			fmt.Println("El stack esta lleno?", s.IsFull())
		}
	}
}
