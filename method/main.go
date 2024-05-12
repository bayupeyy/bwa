package main

import "fmt"

//Ini adalah contoh Struct digunakan di struct yang lain
type Person struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Email     string
	isActive  bool
}

type Group struct {
	Name        string
	Admin       Person
	users       []Person //Logikanya setiap admin mempunyai users dan banyak anggotanya maka membutuhkan user dan slice dari struct person
	IsAvailable bool
}

func (person Person) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", person.FirstName, person.LastName, person.Email)

}

func main() {
	person := Person{1, "Pamungkas", "Aji", 25, "mochamadbayuajie@gmail.com", true}

	result := person.display()

	fmt.Println(result)

	person2 := Person{2, "Link", "Jono", 30, "ajie@gmail.com", true}
	fmt.Println(person2.display())

	// persons := []Person{person, person2}

	// group := Group{"Gamer", person, persons, true}

	// displayGroup(group)

}
