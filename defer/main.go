/*
defer di Golang digunakan untuk menunda eksekusi fungsi atau pernyataan sampai fungsi yang melingkupinya selesai dieksekusi.
Ini berguna untuk memastikan bahwa sumber daya seperti file atau koneksi jaringan ditutup dengan benar, bahkan jika terjadi kesalahan di dalam fungsi.

Berikut adalah contoh penggunaan defer dalam berbagai konteks:

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// Membuka file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Menutup file menggunakan defer
	defer file.Close()

	// Melakukan operasi pada file
	// ...
	fmt.Println("File operations are done")
}
