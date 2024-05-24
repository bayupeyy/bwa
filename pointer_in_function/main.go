package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) { //Jika data ingin di manupulasi maka kita butuh pointer supaya Country bisa berubah contoh awal Address di beri tanda *Address
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  ""}
	ChangeAddressToIndonesia(&alamat) //Alamat membutuhkan Pointer maka beri tanda &

	fmt.Println(alamat)
}

//Materi link reference https://www.youtube.com/watch?v=13skDy7BQyU
