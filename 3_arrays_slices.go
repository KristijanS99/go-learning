/*
Arrays are collection of items that must be of same type
Fixed size, must be known at compile time
Zero index access, len() function, copies original array when assigning to another

Slices are backed by array
Can be created using make() function
len(), cap() and append() functions
Copies refer to same underlying data!!
*/

package main

import "log"

func main() {
	ratings := [3]int{1, 4, 5}
	log.Println("Ratings array:", ratings)
	ratingsCopy := ratings
	ratingsCopy[0] = 111
	log.Println("Ratings first value:", ratings[0],
		"Copy of Ratings first value:", ratingsCopy[0]) // Assigning array creates a copy of it in Go

	ratingsAnotherDef := [...]int{1, 4, 5} // Must know size at compile time
	log.Println(ratingsAnotherDef)

	var ratingsNoElements [3]string
	log.Println("Array with no elements:", ratingsNoElements)
	ratingsNoElements[0] = "Rating test"
	log.Println(ratingsNoElements, "Array length:", len(ratingsNoElements))

	slice := []int{1, 2, 3}
	log.Println(slice, "Length:", len(slice), "Capacity:", cap(slice)) // Slices have length and capacity
	copySlice := slice
	copySlice[0] = 10
	log.Println("Slice and slice copy values:", slice[0], copySlice[0]) // Slices point to the same data, not copied
	spliced := slice[:2]                                                // Slice of the first 2 elements
	log.Println("Value of spliced slice:", spliced)
	spliced[0] = 100
	log.Println("Spliced slice and original:", spliced, slice) // Still points to same underlying data

	sliceUsingMake := make([]int, 5, 100)
	log.Println(sliceUsingMake, "Length:", len(sliceUsingMake), "Capacity:", cap(sliceUsingMake))
	sliceUsingMake = append(sliceUsingMake, 1)
	log.Println("Slice sliceUsingMake after appending 1:", sliceUsingMake)

	sliceUsingMake = append(sliceUsingMake, []int{2, 3, 4}...)
	log.Println("Slice sliceUsingMake after appending {1,2,3}:", sliceUsingMake)
}
