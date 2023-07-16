package main

import (
	"fmt"
	"runtime"

	"github.com/emersonlara/godesdeo/exercicios"
	"github.com/emersonlara/godesdeo/teclado"
	"github.com/emersonlara/godesdeo/variables"
)

func main() {
	//	fmt.Println("Hello world")
	//	variables.MostrarInteiros()
	//variables.RestoVariaveis()

	estado, texto := variables.ConverterTexto(1500)
	fmt.Println(estado)
	fmt.Println(texto)

	valor, frase := exercicios.Devolve("101")

	fmt.Println(valor)
	fmt.Println(frase)

	//	os := runtime.GOOS

	//if os == "Linux."{

	//}else {

	//	}

	if os := runtime.GOOS; os == "Linux." || os == "Ubuntu." {
		fmt.Println("Sistema : ", os)
	} else {
		fmt.Println("Sistema : ", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("E um linux")
	case "darkwin":
		fmt.Println("E um darkwin")
	default:
		fmt.Printf("%s \n", os)
	}

	teclado.IngressoNumeros()

}
