/*
Struct method adalah metode yang terikat pada tipe data struct dalam bahasa Go.
Ini memungkinkan kita untuk mendefinisikan fungsi yang khusus bekerja pada instance dari struct tersebut.
Berikut adalah contoh lengkap bagaimana menggunakan struct method di Go:

*/

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Printf("Hello %s, my name is %s\n", name, customer.Name)
}

func main() {
	daftarCustomer := Customer{
		Name:    "Bayu",
		Address: "Indonesia",
		Age:     25}

	daftarCustomer.sayHello("Aji")
}
