package library

import (
	"fmt"
	"math"
	"strconv"
)

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return (2 * math.Pi * math.Pow(t.JariJari, 2)) + (2 * math.Pi * t.JariJari * t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang) * float64(b.Lebar) * float64(b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * (float64(b.Panjang)*float64(b.Lebar) + float64(b.Panjang)*float64(b.Tinggi) + float64(b.Lebar)*float64(b.Tinggi))
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type PhoneData interface {
	PrintPhoneData()
}

func (p Phone) PrintPhoneData() {
	fmt.Println("name : ", p.Name)
	fmt.Println("brand : ", p.Brand)
	fmt.Println("year : ", p.Year)
	fmt.Print("colors : ")
	for i, v := range p.Colors {
		fmt.Print(v)
		if i+1 < len(p.Colors) {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func LuasPersegi(number int, hasDescription bool) (sentence interface{}) {
	switch {
	case number != 0 && hasDescription:
		sentence = "luas persegi dengan sisi " + strconv.Itoa(number) + " cm adalah " + strconv.Itoa(number*number) + " cm"
	case number != 0 && !hasDescription:
		sentence = number * number
	case number == 0 && hasDescription:
		sentence = "Maaf anda belum menginput sisi dari persegi"
	default:
		sentence = nil
	}
	return
}
