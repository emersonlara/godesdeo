package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nome string
var Estado bool
var Salario float32
var DataFechamento time.Time

func RestoVariaveis() {

	Nome = "Emerson"
	Estado = true
	Salario = 15000
	DataFechamento = time.Now()

	fmt.Println(Nome)
	fmt.Println(Estado)
	fmt.Println(Salario)
	fmt.Println(DataFechamento)

	//MostrarInteiros()
}

func ConverterTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero) //converte para alfa numeroci
	return true, texto
}
