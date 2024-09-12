package main

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) person {
	var newPerson person

	newPerson.Age = age
	newPerson.LastName = lastName
	newPerson.FirstName = firstName

	return newPerson
}

func MakePersonPointer(firstName string, lastName string, age int) *person {
	var newPersonPointer person

	newPersonPointer.Age = age
	newPersonPointer.LastName = lastName
	newPersonPointer.FirstName = firstName

	personPointer := &newPersonPointer

	return personPointer
}
