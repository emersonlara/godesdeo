package main

import (
	"fmt"

	"github.com/emersonlara/godesdeo/variables"
)

func main() {
	//	fmt.Println("Hello world")
	//	variables.MostrarInteiros()
	//variables.RestoVariaveis()
	estado, texto := variables.ConverterTexto(1500)

	fmt.Println(estado)
	fmt.Println(texto)
}
