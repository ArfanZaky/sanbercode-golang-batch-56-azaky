package main

import (
	"fmt"
	"strings"
)

func main() {

	// soal 1

	var bangunDatar hitungBangunDatar
	bangunDatar = segitigaSamaSisi{10, 9}
	bangunDatar = persegiPanjang{7, 8}
	var bangunRuang hitungBangunRuang
	bangunRuang = tabung{10, 10}
	bangunRuang = balok{4, 5, 6}

	// jawaban soal 1
	fmt.Println("===== Segitiga Sama Sisi =====")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	fmt.Println("===== Persegi Panjang =====")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	fmt.Println("===== Tabung =====")
	fmt.Println("volume		:", bangunRuang.volume())
	fmt.Println("luas permukaan	:", bangunRuang.luasPermukaan())

	fmt.Println("===== Balok =====")
	fmt.Println("volume		:", bangunRuang.volume())
	fmt.Println("luas permukaan	:", bangunRuang.luasPermukaan())

	fmt.Println("====================================")
	// soal 2
	var iphone phoneData = phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}

	// jawaban soal 2
	iphone.printPhoneData()

	fmt.Println("====================================")
	// soal 3
	// jawaban soal 3

	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	fmt.Println("====================================")

	// soal 4

	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}
	// jawaban soal 4

	castingKumpulanAngka := append(kumpulanAngkaPertama.([]int), kumpulanAngkaKedua.([]int)...)
	sum := 0
	for _, v := range castingKumpulanAngka {
		sum += v
	}
	fmt.Println(prefix, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(castingKumpulanAngka)), " + "), "[]"), " = ", sum)
	fmt.Println("====================================")

}
