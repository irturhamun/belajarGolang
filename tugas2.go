package main

import "fmt"

func tugas2() {
	//Cek Ganjil/Genap
	var nilai = 12414132
	//Seleksi Kondisi menggunakan IF/ELSE Statement
	if nilai%2 == 1 {
		fmt.Println("Ganjil")
	} else if nilai%2 == 0 {
		fmt.Println("Genap")
	}
}
