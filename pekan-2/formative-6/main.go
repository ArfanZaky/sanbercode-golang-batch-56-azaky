package main

import (
	"fmt"
	"math"
	"strings"
)

func hitungLuasLingkaran(luasLingkaran *float64, jariJari float64) {
	*luasLingkaran = math.Pi * math.Pow(jariJari, 2)
}

func hitungKelilingLingkaran(kelilingLingkaran *float64, jariJari float64) {
	*kelilingLingkaran = 2 * math.Pi * jariJari
}

func introduce(sentence *string, name, gender, job, age string) {
	var honorific string
	if gender == "laki-laki" {
		honorific = "Pak"
	} else if gender == "perempuan" {
		honorific = "Bu"
	} else {
		honorific = "Pak/Bu"
	}
	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", honorific, name, job, age)
}

func buahFavorit(nama string, buah ...string) string {
	return fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah \"%s\"", nama, strings.Join(buah, "\", \""))
}

func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	*dataFilm = append(*dataFilm, map[string]string{"title": title, "duration": duration, "genre": genre, "year": year})
}

func main() {
	// soal 1

	// Tulislah kode seperti di bawah ini, buatlah

	// lalu buatlah function yang nantinya akan memperbarui value dari luas lingkaran dan keliling lingkaran dengan
	// memasuan parameter yaitu pointer luas lingkaran, pointer keliling lingkaran dan jari-jari (wajib menggunakan pointer pada parameter function tersebut)
	var luasLingkaran float64
	var kelilingLingkaran float64
	// jawaban soal 1

	var jariJari float64 = 10 // dummy data

	hitungLuasLingkaran(&luasLingkaran, jariJari)
	hitungKelilingLingkaran(&kelilingLingkaran, jariJari)

	fmt.Println("Luas Lingkaran = ", luasLingkaran)
	fmt.Println("Keliling Lingkaran = ", kelilingLingkaran)

	fmt.Println("====================================")
	// soal 2
	// Tulislah sebuah function dengan nama introduce. pastikan semua parameter pada function introduce di gunakan semuanya (wajib menggunakan pointer)

	// jawaban soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	fmt.Println("====================================")
	// soal 3
	buah := []string{"semangka", "jeruk", "melon", "pepaya"}

	// jawaban soal 3
	buahFavoritJohn := buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)
	fmt.Println("====================================")

	// soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// jawaban soal 4

	// isi dengan jawaban anda untuk menampilkan data
	for indeks, film := range dataFilm {
		fmt.Print(indeks+1, ".")
		for key, value := range film {
			fmt.Println("\t", key, " : ", value)
		}
		fmt.Println()
	}

	fmt.Println("====================================")
}
