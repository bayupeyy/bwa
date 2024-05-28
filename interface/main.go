package main

import "fmt"

type Hewan interface {
	GetNama() string
}
type HasName interface {
	GetName() string
	Age() int
}

//Membuat Func Untuk Animal ( Hewan )
func TampilkanHewan(value Hewan) {
	fmt.Println("Nama Hewan", value.GetNama())
}

//Contoh Membuat Func untuk umur
func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName(), "Umur", value.Age())
}

type Person struct {
	Name string
	Umur int
}

func (umur Person) Age() int {
	return umur.Umur
}

func (person Person) GetName() string {
	return person.Name

}

//Membuat Jadi Beberapa Interface
type Animal struct {
	Name string
}

func (animal Animal) GetNama() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Eko", Umur: 25}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	TampilkanHewan(animal)
}

//Materi Interface https://www.youtube.com/watch?v=IO_vkyJnMas&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ
