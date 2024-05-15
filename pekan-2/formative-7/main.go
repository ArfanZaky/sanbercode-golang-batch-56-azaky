package main

import (
	"fmt"
	"strings"
)

func main() {

	// soal 1
	nanas := buah{"Nanas", "Kuning", false, 9000}
	jeruk := buah{"Jeruk", "Oranye", true, 8000}
	semangka := buah{"Semangka", "Hijau & Merah", true, 10000}
	pisang := buah{"Pisang", "Kuning", false, 5000}

	// jawaban soal 1
	fmt.Printf("Nama\t\tWarna\t\tAda Bijinya\tHarga\n")
	nanas.buah()
	jeruk.buah()
	semangka.buah()
	pisang.buah()

	// soal 2

	Segitiga := segitiga{alas: 3, tinggi: 7}
	Persegi := persegi{sisi: 5}
	PersegiPanjang := persegiPanjang{panjang: 8, lebar: 9}

	// jawaban soal 2
	fmt.Println("Segitiga", "\t: ", Segitiga)
	fmt.Println("Luas Segitiga", "\t: ", Segitiga.luasSegitiga())

	fmt.Println("Persegi", "\t: ", Persegi)
	fmt.Println("Luas Persegi", "\t: ", Persegi.luasPersegi())

	fmt.Println("Persegi Panjang", "\t: ", PersegiPanjang)
	fmt.Println("Luas Persegi Panjang", "\t: ", PersegiPanjang.luasPersegiPanjang())

	// soal 3

	iphone := phone{name: "iPhone SE", brand: "Apple", year: 2016, colors: []string{"Gold"}}
	// jawaban soal 3

	iphone.addColor("Silver")
	fmt.Println(iphone)

	// soal 4

	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)
	// jawaban soal 4
	for indeks, film := range dataFilm {
		fmt.Print(indeks+1, ".\t")
		film.duration /= 60
		s := fmt.Sprintf("%+v\n", film)
		r := strings.NewReplacer("{", "", "}", "", " ", "\n\t", ":", " : ")
		fmt.Println(r.Replace(s))
	}

}
