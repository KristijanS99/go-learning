// All variables must be used
// Lower case scoped to package, uppercase exposed outside
// camelCase preferred, capitalize acronyms (ex. URL)
// Use 'strcov' package for strings conversion
package main

import (
	"log"
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

	log.Printf(":\n %v, %T\n %v, %T\n %v, %T\n %v, %T\n %v, %T\n %v, %T\n",
		i, i, j, j, k, k, theString, theString, strconv.Itoa(i), strconv.Itoa(i), shadowed, shadowed)
}
