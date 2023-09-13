package main

import "fmt"

func main() {
	/// ketika value int32 melebihi int16 nilai diubah menjadi limit max terkecil

	var nilai32 int32 = 12111111
	var nilai64 int64 = int64(nilai32)
	var nilai18 int16 = int16(nilai32)

	fmt.Println(nilai64)
	fmt.Println(nilai18)

	///konversi byte to string

	var nama string = "David"
	var e byte = nama[1]
	var eString string = string(e)

	fmt.Println(eString)
}
