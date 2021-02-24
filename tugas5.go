package main

import (
	"fmt"
	"math/rand"
)

type pelajar1 struct {
	nama string
	umur int
}

func tugas5() {
	var n1, n2, n3 pelajar1
	n1.nama = "Ahmad"
	n1.umur = rand.Intn(100)
	n2.nama = "Abi"
	n2.umur = rand.Intn(100)
	n3.nama = "Zaka"
	n3.umur = rand.Intn(100)
	fmt.Println(n1.nama, n1.umur)
	fmt.Println(n2.nama, n2.umur)
	fmt.Println(n3.nama, n3.umur)

}
