package main

import (
	"fmt"
)

func main() {
	leagueTitles := make(map[string]int)
	leagueTitles["Suderland"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2 := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Println(leagueTitles)
	fmt.Println(recentHead2)

	testMap := map[string]int{"A": 1, "B": 2, "C": 3}

	for key, value := range testMap {
		fmt.Printf("\n %v => %v", key, value)
	}

	testMap["D"] = 4
	fmt.Println(testMap)
	fmt.Println(testMap["C"])
	testMap["A"] = 100
	fmt.Println(testMap)
	delete(testMap, "D")
	fmt.Println(testMap)
}
