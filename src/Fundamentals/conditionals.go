package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	firstRank := "39"
	secondRank := "39"

	if firstRank < secondRank {
		fmt.Println("\nfirst course better" +
			" than second")
	} else if firstRank > secondRank {
		fmt.Println("\nfirst course worse" +
			" than second")
	} else {
		fmt.Println("\nequal value!")
	}

	if firstRank2, secondRank2 := 39, 614; firstRank2 < secondRank2 { // variables here are scoped to the current if statement. shotcircuiting
		fmt.Println("\nfirst course better" +
			" than second")
		if secondRank2 > 100 {
			fmt.Println("\nsecondRank2" + " big!")
		}
	} else if firstRank2 > secondRank2 {
		fmt.Println("\nfirst course worse" +
			" than second")
	} else {
		fmt.Println("\nequal value!")
	}

	switch "docker" {
	case "linux":
		fmt.Println("\nlinux courses")
	case "docker":
		fmt.Println("\ndocker courses")
	case "windows":
		fmt.Println("\nwin courses")
	default:
		fmt.Println("\nNo Match!")
	}

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("even", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("odd", tmpNum)
	}

	_, err := os.Open("C:\\Dan\\Code\\Personal\\Go\\src\\Fundamentals\\conditionals2.go") //C:\\Dan\\Code\\Personal\\Go\\src\\Fundamentals\\conditionals.go
	if err != nil {
		fmt.Println("Error returned is", err)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
