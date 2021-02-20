package main

import "fmt"

func tugas3() {
	//Inisiasi slice (Dynamic Array)
	var buah = []string{"apel", "jeruk", "anggur", "melon"}
	buah = append(buah, "pepaya") //menambah element slice
	for i := 0; i < len(buah); i++ {
		fmt.Println("Element ke - :", i, buah[i])
	}
}
