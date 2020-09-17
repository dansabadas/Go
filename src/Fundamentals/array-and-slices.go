package main

import (
	"fmt"
)

func main() {
	//declare a slice called myCourses
	myCourses := make([]string, 5, 10)
	fmt.Printf("Length is %d. \nCapacity is: %d",
		len(myCourses), cap(myCourses))
	myCourses2 := []string{"course1", " course 2"}
	fmt.Printf("Length is %d. \nCapacity is: %d", len(myCourses2), cap(myCourses2))

	mySlice := []int{1, 2, 3, 40, 55, 666}
	fmt.Println(mySlice[4])
	mySlice[1] = 0
	fmt.Println(mySlice)
	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice)

	mySlice3 := make([]int, 1, 4)
	fmt.Printf("Length is %d. \nCapacity is: %d", len(mySlice3), cap(mySlice3))
	fmt.Println(mySlice3)

	for i := 1; i < 17; i++ {
		mySlice3 = append(mySlice3, i)
		fmt.Printf("Length is %d. \nCapacity is: %d", len(mySlice3), cap(mySlice3))

	}

	mySlice4 := []int{1, 2, 3, 4, 5}

	for _, i := range mySlice4 {
		fmt.Println("for range", i)
	}
	newSlice4 := []int{10, 20}
	mySlice4 = append(mySlice4, newSlice4...)
	fmt.Println(mySlice4)
}
