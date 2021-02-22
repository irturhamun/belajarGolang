package main

import "fmt"

func mapgolang() {
	//inisiasi map di goolang
	var hargamakanan = map[string]int{"ayam_goreng": 15000, "ayam_geprek": 20000}
	//pemanggilan map
	fmt.Println("ayam goreng", hargamakanan["ayam_goreng"])
	fmt.Println("ayam geprek", hargamakanan["ayam_geprek"])
}
