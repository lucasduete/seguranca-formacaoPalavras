package main

import (
	"fmt"
	"strings"
)

var consoantes = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "x", "w", "y", "z"}
var vogais = []string{"a", "e", "i", "o", "u"}

func main() {

	//formWords()
	formWordsWithoutRepetition()
}

func formWords() {
	//Questao3 - Construa um algorítmo que produza todas as palavra de 4
	//letras formadas por uma consoante seguinda de uma vogal.

	//Esta variavel apenas conta o número de palavras possíveis
	var count int = 0

	//Percorre os arrays da seguinte forma: consoante - vogal - consoante - vogal
	//Ao percorre os arrays descarta-se a chave e recupera-se apenas os valores
	for _, letra1 := range consoantes {
		for _, letra2 := range vogais {
			for _, letra3 := range consoantes {
				for _, letra4 := range vogais {
					//Incrementa contador
					count++

					//Imprime a palavra
					fmt.Printf("%s%s%s%s\n", letra1, letra2, letra3, letra4)
				}
			}
		}
	}
	//Imprime o contador
	fmt.Printf("%d palavras possíveis", count)
}

func formWordsWithoutRepetition() {
	//Questao4 - Construa um algorítmo que produza todas as palavra de 6
	//letras formadas por uma consoante seguinda de uma vogal,
	//mas que cada sílaba (consoante+vogal) não seja repetida.

	//Esta variavel apenas conta o número de palavras possíveis
	var count int = 0

	//Percorre os arrays da seguinte forma:
	//	consoante - vogal - consoante - vogal - cosoante -vogal
	//Ao percorre os arrays descarta-se a chave e recupera-se apenas os valores
	for _, letra1 := range consoantes {
		for _, letra2 := range vogais {

			//Definisse a primeira sílaba
			primeiraSilaba := letra1 + letra2

			for _, letra3 := range consoantes {
				for _, letra4 := range vogais {

					//Definisse a segunda sílaba
					segundaSilaba := letra3 + letra4

					//Se a primeira sílaba for diferente da segunda sílaba
					//entao definisse a terceira sílaba
					if strings.Compare(primeiraSilaba, segundaSilaba) != 0 {
						for _, letra5 := range consoantes {
							for _, letra6 := range vogais {

								//Definisse a terceira sílaba
								terceiraSilaba := letra5 + letra6

								//Se todas as silabas forem diferentes então
								//imprime a palavra formada
								if strings.Compare(terceiraSilaba, primeiraSilaba) != 0 &&
									strings.Compare(terceiraSilaba, segundaSilaba) != 0 {
										//Incrementa o contador de palavras possíveis
										count++

										//Imprime a palavra formada
										fmt.Printf("%s%s%s\n", primeiraSilaba, segundaSilaba, terceiraSilaba)
								}
							}
						}
					}

				}
			}
		}
	}
	//Imprime a quantidade de palavras possíveis
	fmt.Printf("%d palavras possíveis", count)
}
