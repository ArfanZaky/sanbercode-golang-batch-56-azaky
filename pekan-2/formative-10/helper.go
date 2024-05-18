package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func logging(sentence string, year int) {
	fmt.Println(sentence, year)
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Panic dengan message: ", message)
	}
}

func kelilingSegitigaSamaSisi(length int, hasDescription bool) (result string, err error) {
	defer endApp()
	switch {
	case length != 0 && hasDescription:
		result = ("keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(length) + " cm adalah " + strconv.Itoa(length*3) + " cm")
	case length != 0 && !hasDescription:
		result = strconv.Itoa(length * 3)
	case length == 0 && hasDescription:
		err = errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	default: // length == 0 && !hasDescription:
		err = errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		panic("ERROR:  sisi = 0 dan boolean = false")
	}
	return
}

func tambahAngka(angkaTambahan int, angka *int) {
	*angka += angkaTambahan
}

func cetakAngka(angka *int) {
	fmt.Println(*angka)
}

func addPhones(phones *[]string, phoneTambahan []string) {
	*phones = append(*phones, phoneTambahan...)
}

func printPhones(phones []string) {
	for i, v := range phones {
		fmt.Println(strconv.Itoa(i+1) + ". " + v)
		time.Sleep(time.Second)
	}
}

func newPrintPhones(newPhones []string, wg *sync.WaitGroup) {
	for i, v := range newPhones {
		fmt.Println(strconv.Itoa(i+1) + ". " + v)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func getMovies(moviesChannel chan string, movies ...string) {
	for _, v := range movies {
		moviesChannel <- v
	}
	close(moviesChannel)
}
