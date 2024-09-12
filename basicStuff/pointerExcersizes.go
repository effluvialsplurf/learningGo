package main

import "fmt"

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

func UpdateSlice(strSlice []string, str string) {
	strSlice[len(strSlice)-1] = str

	fmt.Println(strSlice)
}

func GrowSlice(strSlice []string, str string) {
	strSlice = append(strSlice, str)

	fmt.Println(strSlice)
}

// exploring garbage collection
func MakePeople() {
	for i := 0; i < 10_000_000; i++ {
		guy := person{
			Age:       i,
			LastName:  string(i),
			FirstName: string(i),
		}
		fmt.Println(guy)
	}
}
