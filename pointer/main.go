package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 //Butuh tanda & untuk pass by reference atau mengikuti referensinya bukan Nilainya

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
}
