package main

import "fmt"

func main() {
	var nomor int = 10
	var kata string = "Ini Stringnya"
	var alamatNomor *int = &nomor //pointer ke hard disk
	var alamatkata *string = &kata
	fmt.Println("Nilai", nomor)
	fmt.Println("Alamat Pointer", alamatNomor)
	fmt.Println(kata)
	fmt.Println("Alamat Kata", alamatkata)
}
