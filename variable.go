package main

import "fmt"

func main() {
	var nama string

	nama = "david"
	fmt.Println(nama)

	var NamaKamu = "davidcr"
	fmt.Println(NamaKamu)

	var age = 18
	fmt.Println(age)

	//TODO: Deklarasi var: tidak wajib diganti menjadi := hanya di awal

	negara := "Indonesia"
	fmt.Println(negara)

	negara = "jepang"
	fmt.Println(negara)

	var (
		namaAwal    = "david"
		namaTerakir = "davidcr"
	)
	fmt.Println(namaAwal)
	fmt.Println(namaTerakir)
}
