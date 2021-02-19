package main

import "fmt"

func array() {
	//Inisiasi Array
	var buah = [4]string{"Pepaya", "Mangga", "Pisang", "Jambu"}
	fmt.Println("Jumlah Element:", len(buah))
	fmt.Println("Isi Element:", buah)
	buah[2] = "Anggur" //Ubah buah Pisang jadi Anggur
	fmt.Println("Jumlah Element:", len(buah))
	fmt.Println("Isi Element:", buah)
}
