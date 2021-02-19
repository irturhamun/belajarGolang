package main

import "fmt"

func looping() {
	//For Looping
	for i := 0; i <= 10; i++ {
		fmt.Println("Perulangan for ke", i, "Cek")
	}
	//For looping with variable
	var a = 0
	for a < 10 {
		fmt.Println("Perulangan for dengan Variable ", a)
		a++
	}
	// For loop without initial condition
	for {
		fmt.Println("Perulangan A", a)
		a--
		if a == 0 {
			break
		}
	}
}
