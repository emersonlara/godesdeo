package files

import (
	//"bufio"
	"fmt"
	//	"ioutil"
	"os"

	"github.com/emersonlara/godesdeo/exercicios"
)

var fileName string = "./files/txt/tabela.txt"

func GravaTabudada() {
	var texto string = exercicios.TabelaMultiplica2cont()
	arquivo, err := os.Create("fileName")
	if err != nil {
		fmt.Println("Erro ao criar arquivo" + err.Error())
		return
	}
	fmt.Fprintln(arquivo, texto)
	arquivo.Close()

}

func SomaTabela() {
	var texto string = exercicios.TabelaMultiplica2cont()
	if Append(fileName, texto) == false {
		fmt.Println("Erro ao concactenar conteudo!")

	}
}

func Append(fileName string, texto string) bool {
	arquivo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Erro durante o Append")
		return false
	}
	_, err = arquivo.WriteString(texto)
	if err != nil {
		fmt.Println("Erro durante o WritString")
		return false
	}
	arquivo.Close()
	return true
}
