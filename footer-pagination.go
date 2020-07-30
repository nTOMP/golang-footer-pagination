package main

import (
	"fmt"
	"strconv"
)

/*FooterPagination : Esta função utiliza os parâmetros:
	currentPage => página atual
	totalPages => total de páginas
	boundaries => quantidade de páginas no início e no fim da lista a terem seu número exibido
	around => quantidade de páginas antes e depois da atual a terem seu número exibido
para gerar como resultado a lista com os números de páginas e serem exibidos no rodapé, seguindo o padrão:
	1 2 ... 6 7 8 9 10 ... 19 20
*/
func FooterPagination(currentPage, totalPages, boundaries, around int) string {
	var initBoundaries, initAround, endAround, endBoundaries int
	//controla se as reticências já foram preenchidas para a região
	var printed bool
	var ret string

	//Validação de valores para parâmetros recebidos.
	if currentPage <= 0 {
		fmt.Println("Valor inválido para página atual: O número deve ser maior que 0")
		return ""
	}
	if currentPage > totalPages {
		fmt.Println("Valor inválido para página atual: O número deve ser menor que o total de páginas")
		return ""
	}
	if totalPages <= 0 {
		fmt.Println("Valor inválido para total de páginas: O número deve ser maior que 0")
		return ""
	}
	if boundaries < 0 {
		fmt.Println("Valor inválido para limites: O número deve ser positivo")
		return ""
	}
	if around < 0 {
		fmt.Println("Valor inválido para proximidade: O número deve ser positivo")
		return ""
	}

	//cálculo da região de números a serem exibidos no início da lista
	initBoundaries = boundaries
	//cálculo da região de números a serem exibidos no fim da lista
	endBoundaries = totalPages - boundaries
	//cálculo do início da região de números a serem exibidos próximos da página atual
	initAround = currentPage - around
	//cálculo do fim da região de números a serem exibidos próximos da página atual
	endAround = currentPage + around
	//Para cada número da lista de páginas
	for i := 1; i <= totalPages; i++ {
		//Se o número estiver na região inicial, imprima
		if i <= initBoundaries {
			if i < totalPages {
				fmt.Print(i, " ")
				ret = ret + strconv.Itoa(i) + " "
			} else {
				fmt.Print(i)
				ret = ret + strconv.Itoa(i)
			}
			//As reticências após esta região não foi impressa ainda
			printed = false
			//Se o número estiver na região próxima da página atual, imprima
		} else if i >= initAround && i <= endAround {
			if i < totalPages {
				fmt.Print(i, " ")
				ret = ret + strconv.Itoa(i) + " "
			} else {
				fmt.Print(i)
				ret = ret + strconv.Itoa(i)
			}
			//As reticências após esta região não foi impressa ainda
			printed = false
			//Se o número estiver na região final, imprima
		} else if i > endBoundaries {
			if i < totalPages {
				fmt.Print(i, " ")
				ret = ret + strconv.Itoa(i) + " "
			} else {
				fmt.Print(i)
				ret = ret + strconv.Itoa(i)
			}
			//Para números fora das regiões inicial, próxima ou final, imprima as reticências uma vez
		} else if !printed {
			fmt.Print("... ")
			ret = ret + "... "
			printed = true
		}
	}
	//Quebra de linha para legibilidade
	fmt.Println()
	return ret
}
func main() {
	FooterPagination(4, 5, 1, 0)
	FooterPagination(4, 10, 2, 2)
}
