package main

import (
	"fmt"
	"math"
	"strconv"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return (2 * math.Pi * math.Pow(t.jariJari, 2)) + (2 * math.Pi * t.jariJari * t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * (float64(b.panjang)*float64(b.lebar) + float64(b.panjang)*float64(b.tinggi) + float64(b.lebar)*float64(b.tinggi))
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phoneData interface {
	printPhoneData()
}

func (p phone) printPhoneData() {
	fmt.Println("name : ", p.name)
	fmt.Println("brand : ", p.brand)
	fmt.Println("year : ", p.year)
	fmt.Print("colors : ")
	for i, v := range p.colors {
		fmt.Print(v)
		if i+1 < len(p.colors) {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func luasPersegi(number int, hasDescription bool) (sentence interface{}) {
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
