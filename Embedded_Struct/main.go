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

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.users))

	fmt.Println("Username : ")
	for _, person := range group.users {
		fmt.Println(person.FirstName, person.isActive, person.LastName)
	}
}

func main() {
	person := Person{1, "Pamungkas", "Aji", 25, "mochamadbayuajie@gmail.com", true}
	person2 := Person{2, "Link", "Jono", 30, "ajie@gmail.com", true}

	persons := []Person{person, person2}

	group := Group{"Gamer", person, persons, true}

	displayGroup(group)

}
