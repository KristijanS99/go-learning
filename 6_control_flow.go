package main

import "log"

func main() {
	rolesMap := map[string]int{
		"Admin": 1,
		"User":  2,
	}

	if role, ok := rolesMap["SuperAdmin"]; ok {
		log.Println("SuperAdmin role has value:", role)
	} else if role, ok := rolesMap["Admin"]; ok {
		log.Println("Admin role has value:", role)
	}

	switch i := 2 + 1; i {
	case 1, 2, 3:
		log.Println("1, 2 or 3")
	case 4:
		log.Println("4")
	default:
		log.Println("default")
	}
}
