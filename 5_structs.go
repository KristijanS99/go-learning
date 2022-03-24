package main

import (
	"log"
	"reflect"
)

type Person struct {
	id        int
	firstName string `required max:"50""`
	lastName  string
}

type Student struct {
	Person
	studentId int
}

func main() {
	aPerson := Person{
		id:        1,
		firstName: "John",
		lastName:  "Doe",
	}
	log.Println(aPerson)
	log.Println("Person first name:", aPerson.firstName)
	field, _ := reflect.TypeOf(Person{}).FieldByName("firstName")
	log.Println("firstName validation rules:", field.Tag)

	personCopy := aPerson
	personCopy.lastName = "Test"
	log.Println("aPerson last name:", aPerson.lastName, "personCopy last name:", personCopy.lastName)

	aStudent := Student{
		Person: Person{
			id:        1,
			firstName: "Joe",
			lastName:  "Doe",
		},
		studentId: 1,
	}
	log.Println("Student info:", aStudent)
}
