package main

import "fmt"

func main() {
	type noHp string
	type atau bool

	var SatuAtauDua atau = true
	var noHpOrang1 noHp = "123456789"
	fmt.Println(noHpOrang1)
	fmt.Println(SatuAtauDua)

	var tiga atau = false
	var hp noHp = "dsa"
	fmt.Println(tiga)
	fmt.Println(hp)
}
