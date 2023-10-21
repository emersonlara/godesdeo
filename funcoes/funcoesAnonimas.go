package funcoes

import "fmt"

func Calculo() {

	soma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(soma(10, 45))

}
