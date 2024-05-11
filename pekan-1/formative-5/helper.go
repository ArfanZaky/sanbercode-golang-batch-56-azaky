package main

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

func introduce(name string, gender string, job string, age string) (introduction string) {
	introduction = getTitle(gender) + " " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	return
}

func getTitle(gender string) string {
	if gender == "laki-laki" {
		return "Pak"
	} else if gender == "perempuan" {
		return "Bu"
	}
	return ""
}

func buahFavorit(name string, buah ...string) (introduction string) {
	introduction = "halo nama saya " + name + " dan buah favorit saya adalah "
	for i, b := range buah {
		if i == len(buah)-1 {
			introduction += "\"" + b + "\""
		} else {
			introduction += "\"" + b + "\", "
		}
	}
	return
}
