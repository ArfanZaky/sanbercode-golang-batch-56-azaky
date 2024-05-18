package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {

	// soal 1

	sentence := "Golang Backend Development"
	year := 2021

	// jawaban soal 1
	defer logging(sentence, year)

	// soal 2
	// buatlah function yang returnnya string dan error, dengan kondisi:

	// jika parameter kedua bernilai true maka tampilkan kalimat (asumsi sisinya 2) "keliling segitiga sama sisinya dengan sisi 2 cm adalah 6 cm"
	// jika parameter kedua bernilai false maka tampilkan hanya hasil string angka saja (misal "6")
	// jika parameter pertama 0 dan parameter kedua bernilai true gunakan error dengan message "Maaf anda belum menginput sisi dari segitiga sama sisi"
	// jika parameter pertama 0 dan parameter kedua bernilai false maka tampilkan error dengan message "Maaf anda belum menginput sisi dari segitiga sama sisi", beserta panic yang sudah di recover
	// jawaban soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	fmt.Println()

	// soal 3
	// buatlah function tambahAngka, lalu tampilkan total angka dengan function cetakAngka yang dipanggil menggunakan defer function yang di panggil duluan di func main.
	angka := 1

	// jawaban soal 3
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	// soal 4
	// buatlah function yang menambahkan data di bawah ini ke variabel phones menggunakan pointer:

	// Xiaomi
	// Asus
	// IPhone
	// Samsung
	// Oppo
	// Realme
	// Vivo
	// lalu urutkan data phones tampilkan satu persatu selama per-detik dan sisipkan angka di depan masing-masing data sehingga menghasilkan output seperti ini:
	var phones = []string{}

	// jawaban soal 4
	addPhones(&phones, []string{"Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo"})

	sort.Strings(phones)

	printPhones(phones)

	fmt.Println()

	// soal 5
	// urutkan data phones tampilkan satu persatu per-detik dan sisipkan angka di depan masing-masing data sehingga menghasilkan output seperti ini (gunakan goroutine dan WaitGroup untuk mengerjakan soal ini):
	// gunakan goroutine dan WaitGroup untuk mengerjakan soal ini

	var newPhones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	// jawaban soal 5
	sort.Strings(newPhones)

	var wg sync.WaitGroup

	wg.Add(1)
	go newPrintPhones(newPhones, &wg)

	wg.Wait()

	fmt.Println()

	// soal 6
	// buatlah function getMovies yang akan mengirim data setiap movie ke moviesChannel, output dari soal ini adalah
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	// jawaban soal 6
	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	fmt.Println()

}
