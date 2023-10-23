package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Hallo Guys")
	fmt.Println("Hello World")
	fmt.Println("Hallo Bro")

	var a int
	fmt.Scanln(&a)
	fmt.Println("Hasil Input a-   ", a)

	var b string
		fmt.Scan(&b)
		fmt.Println("Hasil Input b- ", b)

	var nama string = "jerry young"
	fmt.Println(nama)
	var level int = 2
	fmt.Println(level)
	var umur = 20
	fmt.Println(umur)
	var alamat string = "malang"
	fmt.Println(alamat)
	hp := "12345"
	fmt.Println(hp)
	uang := 100
	fmt.Println(uang)

	hasil, err := strconv.Atoi("1")
	fmt.Println(hasil, err)



	var namaku string
	var umurku uint8
	var adalahBenar bool

	fmt.Println(namaku, umurku, adalahBenar)

	var hasilPenjumlah = 5 * 10
	fmt.Println(hasilPenjumlah)
	hasilPenjumlah = uang + umur
	fmt.Println(hasilPenjumlah)
	fmt.Println("nilai awal umur", umur)
	umur++
	fmt.Println("++ pertama", umur)
	umur++
	fmt.Println("++ kedua", umur)
	umur *= 2
	fmt.Println("Setelah di *=2", umur)
	umur /= 5
	fmt.Println("Setelah di /=5", umur)
}