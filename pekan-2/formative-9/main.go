package main

import (
	"fmt"                 // import library
	"formative-9/library" // import library
	"strings"
)

func main() {

	var bangunRuang library.HitungBangunRuang
	var bangunDatar library.HitungBangunDatar

	bangunDatar = library.SegitigaSamaSisi{Alas: 10, Tinggi: 9}
	bangunDatar = library.PersegiPanjang{Panjang: 7, Lebar: 8}
	bangunRuang = library.Tabung{JariJari: 10, Tinggi: 10}
	bangunRuang = library.Balok{Panjang: 4, Lebar: 5, Tinggi: 6}

	fmt.Println("===== Segitiga Sama Sisi =====")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())

	fmt.Println("===== Persegi Panjang =====")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())

	fmt.Println("===== Tabung =====")
	fmt.Println("volume		:", bangunRuang.Volume())
	fmt.Println("luas permukaan	:", bangunRuang.LuasPermukaan())

	fmt.Println("===== Balok =====")
	fmt.Println("volume		:", bangunRuang.Volume())
	fmt.Println("luas permukaan	:", bangunRuang.LuasPermukaan())

	fmt.Println("====================================")
	var iphone library.PhoneData = library.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	iphone.PrintPhoneData()

	fmt.Println("====================================")

	fmt.Println(library.LuasPersegi(4, true))

	fmt.Println(library.LuasPersegi(8, false))

	fmt.Println(library.LuasPersegi(0, true))

	fmt.Println(library.LuasPersegi(0, false))

	fmt.Println("====================================")

	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	castingKumpulanAngka := append(kumpulanAngkaPertama.([]int), kumpulanAngkaKedua.([]int)...)
	sum := 0
	for _, v := range castingKumpulanAngka {
		sum += v
	}
	fmt.Println(prefix, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(castingKumpulanAngka)), " + "), "[]"), " = ", sum)
	fmt.Println("====================================")

}
