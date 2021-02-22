package main

import "fmt"

func main() {
	var s1 pelajar
	var s2 = pelajar{"Ihlas", 21}
	s1.nama = "Amin"
	s1.umur = 20
	s1.namasaya()
	s2.namasaya()
}

type pelajar struct {
	nama string
	umur int
}

func (s pelajar) namasaya() {
	fmt.Println("Nama saya adalah", s.nama)
	fmt.Println("Umur saya ", s.umur, "Tahun")
}
