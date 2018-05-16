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
			for _, letra3 := range consoantes {
				for _, letra4 := range vogais {
					if !(strings.Compare(letra1, letra3) == 0 && strings.Compare(letra2, letra4) == 0) {
						count++
						fmt.Printf("%s%s%s%s\n", letra1, letra2, letra3, letra4)
					}
				}
			}
		}
	}
	fmt.Printf("%d palavras possíveis", count)
}
