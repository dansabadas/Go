package main

import (
	"fmt"
)

func main() {
	type CourseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	var DockerDeep CourseMeta
	DockerDeep3 := new(CourseMeta) // this gives us a pointer
	DockerDeep2 := CourseMeta{
		Author: "Dan",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeep)
	fmt.Println(DockerDeep3)
	fmt.Println(DockerDeep2)
	fmt.Println("DockerDeep2.Author =", DockerDeep2.Author)
	DockerDeep2.Rating = 2
	fmt.Println(DockerDeep2)
}
