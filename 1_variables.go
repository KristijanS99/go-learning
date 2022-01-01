package main

import (
	"fmt"
	"strconv"
)

var shadowed = 30

func main() {
	var i int
	i = 20
	var j = 30.5
	k := 40
	theString := "The string"
	shadowed := 50

	fmt.Printf(" %v, %T\n %v, %T\n %v, %T\n %v, %T\n %v, %T\n %v, %T\n",
		i, i, j, j, k, k, theString, theString, strconv.Itoa(i), strconv.Itoa(i), shadowed, shadowed)
}
