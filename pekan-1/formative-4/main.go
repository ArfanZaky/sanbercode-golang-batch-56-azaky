package main

import (
	"fmt"
)

func main() {
	// soal 1
	// Pada tugas ini kamu diminta untuk melakukan looping. Untuk membuat tantangan ini lebih menarik, kamu juga diminta untuk memenuhi syarat tertentu yaitu:

	// SYARAT:
	// A. Jika angka ganjil maka tampilkan Santai
	// B. Jika angka genap maka tampilkan Berkualitas
	// C. Jika angka yang sedang ditampilkan adalah kelipatan 3 DAN angka ganjil maka tampilkan I Love Coding.

	// jawaban soal 1
	for i := 1; i <= 20; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "Berkualitas")
		case i%3 == 0 && i%2 != 0:
			fmt.Println(i, "I Love Coding")
		default:
			fmt.Println(i, "Santai")
		}
	}

	// soal 2
	// Kali ini kamu diminta untuk menampilkan sebuah segitiga dengan tanda pagar (#) dengan dimensi tinggi 7 dan alas 7. gunakan looping untuk mengerjakan soal ini
	// jawaban soal 2
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	fmt.Println("====================================")

	// soal 3
	var kalimatSoal3 = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	// jawaban soal 3
	var kalimatSoal3Baru []string
	for i := 2; i < len(kalimatSoal3); i++ {
		kalimatSoal3Baru = append(kalimatSoal3Baru, kalimatSoal3[i])
	}
	fmt.Println(kalimatSoal3Baru)
	fmt.Println("====================================")

	// soal 4
	var sayuran = []string{}
	// jawaban soal 4
	var tambahanSayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}
	sayuran = append(sayuran, tambahanSayuran...)
	for i := 0; i < len(sayuran); i++ {
		fmt.Println(i+1, sayuran[i])
	}

	fmt.Println("====================================")

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	// jawaban soal 5
	for key, value := range satuan {
		fmt.Println(key, ":", value)
	}
	var volumeBalok = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	fmt.Println("Volume Balok :", volumeBalok)
	fmt.Println("====================================")
}
