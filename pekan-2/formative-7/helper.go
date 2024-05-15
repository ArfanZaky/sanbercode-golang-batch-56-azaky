package main

import (
	"fmt"
)

func (b buah) buah() {
	fmt.Printf("%s\t\t%s\t\t", b.nama, b.warna)
	if b.adaBijinya {
		fmt.Printf("Ada")
	} else {
		fmt.Printf("Tidak")
	}
	fmt.Println("\t\t", b.harga)
}

func (s segitiga) luasSegitiga() float64 {
	return float64(s.alas) * float64(s.tinggi) / 2
}

func (p persegi) luasPersegi() int {
	return p.sisi * p.sisi
}

func (pP persegiPanjang) luasPersegiPanjang() int {
	return pP.panjang * pP.lebar
}

func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	*dataFilm = append(*dataFilm, movie{title, duration, genre, year})
}
