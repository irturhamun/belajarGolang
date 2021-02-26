//goroutine mini thread di Golang
package main

import (
	"fmt"
	"runtime"
)

func gorutin1() {
	runtime.GOMAXPROCS(2)
	//Asyncronous Proccess
	go tampilkanpesanku(5, "Ini Pesanku Terkirim")
	tampilkanpesanku(5, "Saya tiba")

	//Blocking untuk menahan process
	var input string
	fmt.Scanln(&input)
}

func tampilkanpesanku(x int, pesan string) {
	for i := 0; i < x; i++ {
		fmt.Println((i + 1), pesan)
	}
}
