/*
Penggunaan recover:
Fungsi recover digunakan untuk menangkap panic dan memungkinkan program untuk melanjutkan eksekusi normal.
recover hanya bekerja jika dipanggil dari dalam fungsi defer.
*/

package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Sebelum panic")
	panic("Terjadi kesalahan fatal!")
	fmt.Println("Setelah panic") // Tidak akan dieksekusi
}
