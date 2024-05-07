package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// soal 1
	// jawaban soal 3
	var Bootcamp string = "Bootcamp"
	var Digital string = "Digital"
	var Skill string = "Skill"
	var Sanbercode string = "Sanbercode"
	var Golang string = "Golang"
	var kalimatSoal1 = fmt.Sprintf("%s %s %s %s %s", Bootcamp, Digital, Skill, Sanbercode, Golang)
	fmt.Println(kalimatSoal1)
	fmt.Println("====================================")

	// soal 2
	var halo string = "Halo Dunia"
	// jawaban soal 2
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(halo)
	fmt.Println("====================================")

	// soal 3
	var kataPertama string = "saya"
	var kataKedua string = "senang"
	var kataKetiga string = "belajar"
	var kataKeempat string = "golang"

	// jawaban soal 3
	var kalimat = fmt.Sprintf("%s %s %s %s", kataPertama, strings.Title(kataKedua), uppercaseLastChar(kataKetiga), strings.ToUpper(kataKeempat))
	fmt.Println(kalimat)
	fmt.Println("====================================")

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var a, _ = strconv.Atoi(angkaPertama)
	var b, _ = strconv.Atoi(angkaKedua)
	var c, _ = strconv.Atoi(angkaKetiga)
	var d, _ = strconv.Atoi(angkaKeempat)
	var hasil = a + b + c + d
	// jawaban soal 4
	fmt.Println(hasil)
	fmt.Println("====================================")

	// soal 5
	kalimattask5 := "halo halo bandung"
	angka := 2021

	kalimattask5 = strings.ReplaceAll(kalimattask5, "halo", "Hi")
	kalimattask5 = fmt.Sprintf("%s - %d", kalimattask5, angka)
	// jawaban soal 5
	fmt.Println(kalimattask5)
	fmt.Println("====================================")
}

func uppercaseLastChar(s string) string {
	chars := []rune(s)

	if len(chars) > 0 {
		chars[len(chars)-1] = unicode.ToUpper(chars[len(chars)-1])
	}

	return string(chars)
}
