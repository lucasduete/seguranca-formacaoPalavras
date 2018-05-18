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
	var count int = 0
	for _, letra1 := range consoantes {
		for _, letra2 := range vogais {
			for _, letra3 := range consoantes {
				for _, letra4 := range vogais {
					count++
					fmt.Printf("%s%s%s%s\n", letra1, letra2, letra3, letra4)
				}
			}
		}
	}
	fmt.Printf("%d palavras possíveis", count)
}

func formWordsWithoutRepetition() {
	var count int = 0
	for _, letra1 := range consoantes {
		for _, letra2 := range vogais {

			primeiraSilaba := letra1 + letra2

			for _, letra3 := range consoantes {
				for _, letra4 := range vogais {

					segundaSilaba := letra3 + letra4

					if strings.Compare(primeiraSilaba, segundaSilaba) != 0 {
						for _, letra5 := range consoantes {
							for _, letra6 := range vogais {

								terceiraSilaba := letra5 + letra6

								if strings.Compare(terceiraSilaba, primeiraSilaba) != 0 &&
									strings.Compare(terceiraSilaba, segundaSilaba) != 0 {
										count++
										fmt.Printf("%s%s%s%s%s%s\n", letra1, letra2, letra3, letra4, letra5, letra6)
								}
							}
						}
					}

				}
			}
		}
	}
	fmt.Printf("%d palavras possíveis", count)
}
