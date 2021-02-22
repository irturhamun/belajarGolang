package main

import "fmt"

//definisi struct
type pelajarku struct {
	nama  string
	kelas string
	umur  int
}

func structure() {
	var x1 pelajarku
	x1.nama = "Anthony"
	x1.kelas = "XII-IPA-2"
	x1.umur = 17
	fmt.Println("nama :", x1.nama)
	fmt.Println("kelas :", x1.kelas)
	fmt.Println("umur :", x1.umur)
}
