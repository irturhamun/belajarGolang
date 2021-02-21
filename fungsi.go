package main

import "fmt"

func fungsi() {
	fmt.Println(fungsi_pertama(10, 5))  //Pemanggilan fungsi
	var x1, x2 = fungsi_multiple(20, 5) //pemanggilan fungsi multiple return
	fmt.Println(x1)
	fmt.Println(x2)
}

func fungsi_pertama(x int, y int) int { //ini inisiasi fungsi mengembalikan nilai string
	var z = x + y

	return z
}

//multiple return function

func fungsi_multiple(x int, y int) (int, int) {
	var a = x * y
	var b = x / y
	return a, b
}
