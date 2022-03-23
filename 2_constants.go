/*
All variables must be used
Lower case scoped to package, uppercase exposed outside
camelCase preferred, capitalize acronyms (ex. URL)
Use 'strcov' package for strings conversion
*/
package main

import (
	"log"
)

func main() {
	const firstConstant int = 42
	const secondConstant = 23.2
	log.Printf(": %v, %T", firstConstant, firstConstant)
	log.Printf(": %v, %T", secondConstant, secondConstant)

	const (
		_ = iota
		enumeratedConst1
		enumeratedConst2
		enumeratedConst3
	)
	log.Printf(": %v, %T", enumeratedConst1, enumeratedConst1)
	log.Printf(": %v, %T", enumeratedConst2, enumeratedConst2)
	log.Printf(": %v, %T", enumeratedConst3, enumeratedConst3)

	const (
		isMale     = 1 << iota // 001
		isAdult                // 010
		isEmployed             // 100
	)
	var personInfo byte = isAdult | isEmployed // 110
	log.Printf(": %v, %T", personInfo, personInfo)
	log.Printf(": Is male? %v", personInfo&isMale == isMale)    // 110 & 001 => 000 equals 001 ? -> false
	log.Printf(": Is adult? %v", personInfo&isAdult == isAdult) // 110 & 010 => 010 equals 010 ? -> true
}
