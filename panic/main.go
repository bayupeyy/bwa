/*
Fungsi panic dalam bahasa pemrograman, khususnya di Go (Golang),
digunakan untuk menandakan situasi yang tidak dapat diatasi oleh program,
seperti kesalahan kritis yang mengharuskan program untuk dihentikan.
Berikut adalah penjelasan lebih rinci tentang fungsi panic:

*/

package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups Error")
	}
}

func main() {
	runApp(true)
}
