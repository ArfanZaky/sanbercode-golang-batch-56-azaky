package webserver

import (
	"fmt"
	"net/http"
)

func Volume(r float64, t float64) float64 {
	return 3.14 * r * r * t
}

func LuasAlas(r float64) float64 {
	return 3.14 * r * r
}

func KelilingAlas(r float64) float64 {
	return 2 * 3.14 * r
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "jariJari : 7, tinggi: 10, volume : ", Volume(7, 10), ", luas alas: ", LuasAlas(7), ", keliling alas: ", KelilingAlas(7))
}
