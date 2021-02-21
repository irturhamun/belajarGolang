package main

import "fmt"

func Map() {
	//inisiasi map di goolang
	var harga_makanan = map[string]int{"ayam_goreng": 15000, "ayam_geprek": 20000}
	//pemanggilan map
	fmt.Println("ayam goreng", harga_makanan["ayam_goreng"])
	fmt.Println("ayam geprek", harga_makanan["ayam_geprek"])
}
