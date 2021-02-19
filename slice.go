package main

import "fmt"

func main() {
	//Inisiasi slice (Dynamic Array)
	var buah = []string{"apel", "mangga", "jeruk", "anggur"}
	buah = append(buah, "pisang") //menambah element slice

	fmt.Println(buah)
	fmt.Println("jumlah element :", len(buah))
}
