/*
All variables must be used
Lower case scoped to package, uppercase exposed outside
camelCase preferred, capitalize acronyms (ex. URL)
Use 'strcov' package for strings conversion
*/
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
	byteSlice := []byte(theString)
	shadowed := 50

	log.Printf(": %v, %T", i, i)
	log.Printf(": %v, %T", j, j)
	log.Printf(": %v, %T", k, k)
	log.Printf(": %v, %T", theString, theString)
	log.Printf(": %v, %T", strconv.Itoa(i), strconv.Itoa(i))
	log.Printf(": %v, %T", shadowed, shadowed)
	log.Printf(": %v, %T", byteSlice, byteSlice)
	log.Printf(": %v, %T", string(byteSlice), string(byteSlice))

	var theBoolean bool
	log.Printf(": %v, %T", theBoolean, theBoolean)
	theBoolean = true
	log.Printf(": %v, %T", theBoolean, theBoolean)
}
