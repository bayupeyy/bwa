package main

import (
	"fmt"
)

func main() {
	var age int = 25

	fmt.Println(age)

	tahun := 1

	if tahun > 10 {
		fmt.Printf("Tahun ini boleh dolan")
	} else {
		fmt.Printf("Belum boleh dolan")
	}

	score := 80
	var grade string

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else {
		grade = "D"
	}

	fmt.Println("\nNilai anda adalah ", score, "\nGrade anda adalah ", grade)

	number := 3

	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	default:
		fmt.Println("Nomor bukan 1 atau 2")
	}

	for i := 0; i < 100; i++ {
		fmt.Println("Coding!")
	}

	tittle := "Ini adalah Golang"

	for index, letter := range tittle {
		letterString := string(letter)
		if index%2 == 0 || letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" ||
			letterString == "o" {
			fmt.Println("Ini index", index, "Letter", string(letter))
		}

	}
}
