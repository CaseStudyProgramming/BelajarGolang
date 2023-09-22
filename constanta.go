package main

import "fmt"

func main() {
	const namaAwal string = "david"

	const namaAkhir = "cristian"
	const NamaTengah = "yohoo"

	/// tanpa value constant const tidak dikomplain

	/// ketika di run tidak bisa diubah akan terjadi error

	namaAwal = "yoza"
	namaAkhir = "hatsune"

	const (
		awal  = "akhir"
		akhir = "awal"
	)

	fmt.Println(awal)
	fmt.Println(akhir)
	fmt.Println(NamaTengah)
}
