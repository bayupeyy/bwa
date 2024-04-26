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

	//Ini adalah Array

	iniArray := [...]string{
		"A",
		"B",
		"C",
		"D",
	}

	fmt.Println(iniArray)

	for index, iniArray := range iniArray {
		fmt.Println("Index : ", index, "Array : ", iniArray)
	}

	//Ini adalah Slice

	var jenisHP []string

	jenisHP = append(jenisHP, "Samsung")
	fmt.Println(jenisHP)

	jenisUlo := []string{"Nogo", " Cukrik"}
	jenisUlo = append(jenisUlo)
	fmt.Println(jenisUlo)

	var jenisSepatu []string
	jenisSepatu = append(jenisSepatu, "Nike")
	jenisSepatu = append(jenisSepatu, "Adidas")
	jenisSepatu = append(jenisSepatu, "Puma")
	jenisSepatu = append(jenisSepatu, "Converse")
	jenisSepatu = append(jenisSepatu, "Vans")
	jenisSepatu = append(jenisSepatu, "New Balance")
	jenisSepatu = append(jenisSepatu, "Reebok")
	jenisSepatu = append(jenisSepatu, "Asics")

	for index, jenisSepatu := range jenisSepatu {
		fmt.Println("Index : ", index, "Jenis Sepatu : ", jenisSepatu)

	}
}
