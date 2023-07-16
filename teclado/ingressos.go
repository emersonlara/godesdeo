package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var legenda string
var err error

func IngressoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite o numero 1:")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Erro ao processar numero 1 !" + err.Error())
		}
	}

	fmt.Println("Digite o numero 2:")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Erro ao processar numero  2!" + err.Error())
		}
	}

	fmt.Println("Digite a legenda:")
	if scanner.Scan() {
		legenda = scanner.Text()
	}

	fmt.Println(legenda, numero1*numero2)
}
