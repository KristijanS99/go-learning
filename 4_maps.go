/*
Maps are passed by reference
*/

package main

import "log"

func main() {
	rolesMap := map[string]int{
		"Admin": 1,
		"User":  2,
	}
	log.Println("Roles map:", rolesMap)
	rolesMap["SuperAdmin"] = 3
	log.Println(rolesMap["SuperAdmin"])
	log.Println("Roles map:", rolesMap) // Order not guaranteed after altering

	delete(rolesMap, "SuperAdmin")
	log.Println("Roles map:", rolesMap)

	_, ok := rolesMap["SuperAdmin"]
	log.Println("Role SuperAdmin was found:", ok)
	log.Println("Size of map:", len(rolesMap))
}
