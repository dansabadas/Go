package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 2; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("BOOM!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	coursesInProg := []string{"docker1", "C#2", "Python3"}
	//coursesCompleted := []string{"javascript1", "SQL2", "Rucy on Rails"}
	for _, i := range coursesInProg {
		fmt.Println(i)
	}

	for j, i := range coursesInProg {
		fmt.Println(i, j)
	}

	for timer := 10; timer >= 0; timer-- {
		if timer%2 == 0 {
			continue
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}
