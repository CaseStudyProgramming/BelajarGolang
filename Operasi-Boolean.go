package main

import "fmt"

func main() {
	// Operasi-Boolean
	var nilaiAkhir = 90

	var Absensi = 80

	var lulusJikaNilai bool = nilaiAkhir > 80
	var lulusAbsensi bool = Absensi > 80

	var lulus bool = lulusJikaNilai && lulusAbsensi
	fmt.Println(lulus)
}
