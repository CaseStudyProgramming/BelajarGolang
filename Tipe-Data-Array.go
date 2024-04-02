package main

import "fmt"

func main() {
	var (
		///default jika string kosong atau int kosong adalah "" dan 0
		names [3]string
		///... digunakan jika tidak tau panjang secara pasti
		values = [...]int{
			1, 2, 3,
		}
	)
	names[0] = "David"
	names[1] = "Cristia"
	names[2] = "Putra S"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(values)
	fmt.Println(len(values))
}
