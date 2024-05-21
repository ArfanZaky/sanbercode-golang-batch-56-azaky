package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

// buatlah API POST Nilai Mahasiswa untuk menambahkan data NilaiMahasiswa, ada beberapa ketentuan diantaranya:
// POST dapat menerima data dalam bentuk JSON
// data yang di input hanya boleh Nama, MataKuliah dan Nilai saja, untuk ID dan IndeksNilai harus diolah terlebih dahulu baru di masukkan ke tambahkan ke NilaiMahasiswa
// Nilai hanya boleh diisi maksimal dengan angka 100
// untuk mengisi IndeksNilai memiliki kondisi: Nilai >= 80 indeksnya A, Nilai >= 70 dan Nilai < 80 indeksnya B, Nilai >= 60 dan Nilai < 70, indeksnya c Nilai >= 50 dan Nilai < 60 indeksnya D, Nilai < 50 indeksnya E
// harus memasukkan dulu username dan password baru bisa melakukan POST Nilai Mahasiswa (Menggunakan Basic Auth)
//  buatlah API GET Nilai Mahasiswa untuk menampilkan semua data NilaiMahasiswa

// GetMovies
func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilaiNilaiMahasiswa, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilaiNilaiMahasiswa)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)

}

func PostNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if !basicAuth(w, r) {
		return
	}
	if r.Method == "POST" {
		var nilaiMahasiswa NilaiMahasiswa
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(&nilaiMahasiswa)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			if nilaiMahasiswa.Nama == "" || nilaiMahasiswa.MataKuliah == "" || nilaiMahasiswa.Nilai == 0 {
				http.Error(w, "Nama, MataKuliah, Nilai harus diisi", http.StatusBadRequest)
				return
			}
			if nilaiMahasiswa.Nilai > 100 {
				http.Error(w, "Nilai tidak boleh lebih dari 100", http.StatusBadRequest)
				return
			}

			if nilaiMahasiswa.Nilai >= 80 {
				nilaiMahasiswa.IndeksNilai = "A"
			} else if nilaiMahasiswa.Nilai >= 70 {
				nilaiMahasiswa.IndeksNilai = "B"
			} else if nilaiMahasiswa.Nilai >= 60 {
				nilaiMahasiswa.IndeksNilai = "C"
			} else if nilaiMahasiswa.Nilai >= 50 {
				nilaiMahasiswa.IndeksNilai = "D"
			} else {
				nilaiMahasiswa.IndeksNilai = "E"
			}

			nilaiMahasiswa.ID = uint(len(nilaiNilaiMahasiswa) + 1)
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)

			dataNilaiMahasiswa, err := json.Marshal(nilaiMahasiswa)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(dataNilaiMahasiswa)
			return
		} else {
			http.Error(w, "Content-Type harus application/json", http.StatusBadRequest)
			return
		}

	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func basicAuth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
	if username != "admin" || password != "admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
	return true
}

func main() {
	// tugas 1
	http.HandleFunc("/post-movies", PostNilaiMahasiswa)
	// tugas 2
	http.HandleFunc("/get-movies", getNilaiMahasiswa)

	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
