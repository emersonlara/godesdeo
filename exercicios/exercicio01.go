package exercicios

import (
	"strconv"
)

func Devolve(palavra string) (int, string) {
	var valor int
	var frase string

	valor, err := strconv.Atoi(palavra)
	if err != nil {
		// Lidar com erro, caso a conversão falhe
		frase = "Valor inválido"
		return valor, frase + " pilha de erro : " + err.Error()
	}

	if valor > 100 {
		frase = "Maior que 100"
	} else {
		frase = "Menor que 100"
	}

	return valor, frase
}
