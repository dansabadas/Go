package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "func basics"
	author := "nigel"

	fmt.Println(converter(module, author))

	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16, 5, 6)
	fmt.Println(bestFinish)
}

func converter(s1, author string) (string, s2 string) { //or (str1, s2 string) or (string, string) as return declaration
	s1 = strings.Title(s1)
	author = strings.Title(author)

	return s1, author
}

func bestLeagueFinishes(finishes ...int) int { // this is a variadic function - with any (dynamic) number of parameters
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
