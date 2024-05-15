package main

type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type movie struct {
	title    string
	duration int
	genre    string
	year     int
}
