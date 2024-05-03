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

	//ini adalah Map

	var myMap map[string]int
	myMap = map[string]int{}

	myMap["Ruby"] = 9
	myMap["Jono"] = 10
	myMap["Jono"] = 11

	fmt.Println(myMap)
	//untuk mendapatkan nilai dari salah satu
	fmt.Println(myMap["Ruby"])

	//Bentuk Map lain
	sepatu := map[string]int{
		"Nike":   10,
		"Adidas": 11,
		"Puma":   12,
	}

	for key, value := range sepatu {
		fmt.Println(key, value)
	}

	//untuk menghapus data
	//delete(sepatu, "Nike")

	fmt.Println(sepatu)

	fmt.Println("=================")
	value, isAvailable := myMap["Jonioo"]

	fmt.Println(value, isAvailable)

	//Ini adalah Slice of map
	dataSiswa := []map[string]string{
		{"nama": "Jono", "kelas": "10"},
		{"nama": "Karman", "kelas": "10"},
		{"nama": "Cino", "kelas": "10"},
	}

	fmt.Println(dataSiswa)

	for _, siswa := range dataSiswa {
		fmt.Println(siswa["nama"])

	}

	//Latihan soal
	//Hitung rata rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var total int

	for _, score := range scores {
		total = total + score
	}

	length := len(scores)
	average := float64(total) / float64(length)

	fmt.Println("Rata ratanya adalah", average)

	//Memasukkan Nilai yang =< 90

	var goodScores []int
	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}

		fmt.Println(goodScores)
	}

	//Fungsi
	myResult()

	//Percobaan
	for j := 0; j <= 10; j++ {
		for k := 1; k <= j; k++ {
			fmt.Print("s")
		}
		fmt.Println()
	}

	hasil("belajar")

	sentence := printMyResult("Saya Sedang")
	fmt.Println(sentence)

	//Perkalian
	perkalian := perkalian(9, 10)
	fmt.Println(perkalian)

	fmt.Println("==================================")
	fmt.Println("===============Ini Adalah Multiple Return Function===================")

	//Multiple Return Function

	tambah, kurang, kali, bagi := calculator(10, 10, 10)

	fmt.Println(tambah)
	fmt.Println(kurang)
	fmt.Println(kali)
	fmt.Println(bagi)

}

func myResult() {
	fmt.Println("Belajar Golang")
}

//Menggunakan Paramater

func hasil(iniHasil string) {
	fmt.Println(iniHasil)
}

//Fungsi dengan nilai balik

func printMyResult(sentence string) string {
	newSetence := sentence + " Belajar Golang"
	return newSetence
}

//Func Perkalian

func perkalian(number int, numberTwo int) int {
	result := number * numberTwo
	return result
}

//Multiple Return Func

func calculator(angka1, angka2, angka3 int) (int, int, int, float64) {
	penambahan := angka1 + angka2 + angka3
	pengurangan := angka1 - angka2 - angka3
	kali := angka1 * angka2 * angka3
	pembagian := float64(angka1) / float64(angka2) / float64(angka3)

	return penambahan, pengurangan, kali, pembagian
}
