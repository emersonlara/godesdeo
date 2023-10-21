package iteracoes

import (
	"fmt"
)

func Iterar() {

	i := 0
	for i <= 100 {
		fmt.Println(i)
		i++
	}
}

func Iteracao() {

	for i := 0; i <= 1000; i++ {

		fmt.Println(i)
	}
}

func Iteracao2() {

	for i := 0; i <= 1000; i += 5 {

		fmt.Println(i)
	}
}

func IteracaoInversa() {
	for i := 1000; i > 1; i -= 5 {
		fmt.Println(i)
	}
}

func IteracaoInversa2() {
	for i := 10; i > 1; i-- {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}
