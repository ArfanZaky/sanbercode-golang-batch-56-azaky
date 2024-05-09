package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	// jawaban soal 1
	luasPersegiPanjang = panjang * lebar
	kelilingPersegiPanjang = 2 * (panjang + lebar)
	luasSegitiga = (alas * tinggi) / 2
	fmt.Println(luasPersegiPanjang, kelilingPersegiPanjang, luasSegitiga)
	fmt.Println("====================================")

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50
	var name = []string{"John", "Doe"}
	// jawaban soal 2
	var array = []int{nilaiJohn, nilaiDoe}
	for i, nilai := range array {
		switch {
		case nilai >= 80:
			fmt.Println(name[i], "mendaapatkan nilai A")
		case nilai >= 70 && nilai < 80:
			fmt.Println(name[i], "mendapatkan nilai B")
		case nilai >= 60 && nilai < 70:
			fmt.Println(name[i], "mendapatkan nilai C")
		case nilai >= 50 && nilai < 60:
			fmt.Println(name[i], "mendapatkan nilai D")
		default:
			fmt.Println(name[i], "mendapatkan nilai E")
		}
	}
	fmt.Println("====================================")

	// soal 3
	var tanggal = 22
	var bulan = 12
	var tahun = 1998

	// jawaban soal 3
	switch bulan {
	case 1:
		fmt.Println(tanggal, "Januari", tahun)
	case 2:
		fmt.Println(tanggal, "Februari", tahun)
	case 3:
		fmt.Println(tanggal, "Maret", tahun)
	case 4:
		fmt.Println(tanggal, "April", tahun)
	case 5:
		fmt.Println(tanggal, "Mei", tahun)
	case 6:
		fmt.Println(tanggal, "Juni", tahun)
	case 7:
		fmt.Println(tanggal, "Juli", tahun)
	case 8:
		fmt.Println(tanggal, "Agustus", tahun)
	case 9:
		fmt.Println(tanggal, "September", tahun)
	case 10:
		fmt.Println(tanggal, "Oktober", tahun)
	case 11:
		fmt.Println(tanggal, "November", tahun)
	case 12:
		fmt.Println(tanggal, "Desember", tahun)
	}
	fmt.Println("====================================")

	// soal 4
	// Baby boomer, kelahiran 1944 s.d 1964
	// Generasi X, kelahiran 1965 s.d 1979
	// Generasi Y (Millenials), kelahiran 1980 s.d 1994
	// Generasi Z, kelahiran 1995 s.d 2015
	// jawaban soal 4
	var tahunSaya = 1998
	switch {
	case tahunSaya >= 1944 && tahunSaya <= 1964:
		fmt.Println("Baby boomer")
	case tahunSaya >= 1965 && tahunSaya <= 1979:
		fmt.Println("Generasi X")
	case tahunSaya >= 1980 && tahunSaya <= 1994:
		fmt.Println("Generasi Y (Millenials)")
	case tahunSaya >= 1995 && tahunSaya <= 2015:
		fmt.Println("Generasi Z")
	default:
		fmt.Println("Tahun lahir tidak sesuai dengan kategori")
	}

	fmt.Println("====================================")
}
