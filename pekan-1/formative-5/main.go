package main

import (
	"fmt"
)

func main() {
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	// jawaban soal 1
	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	fmt.Println("====================================")
	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	sarah := introduce("Sarah", "perempuan", "model", "28")
	// jawaban soal 2
	fmt.Println(john)  // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	fmt.Println("====================================")
	// soal 3

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var name = "John"

	var buahFavoritJohn = buahFavorit(name, buah...)

	// jawaban soal 3
	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	fmt.Println("====================================")
	// soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	// jawaban soal 4

	var tambahDataFilm = func(title string, jam string, genre string, tahun string) {
		dataFilm = append(dataFilm, map[string]string{
			"genre": genre,
			"jam":   jam,
			"tahun": tahun,
			"title": title,
		})
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

	fmt.Println("====================================")
}
