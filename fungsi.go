package main

import "fmt"

func fungsi() {
	fmt.Println(fungsipertama(10, 5))  //Pemanggilan fungsi
	var x1, x2 = fungsimultiple(20, 5) //pemanggilan fungsi multiple return
	fmt.Println(x1)
	fmt.Println(x2)
}

func fungsipertama(x int, y int) int { //ini inisiasi fungsi mengembalikan nilai string
	var z = x + y

	return z
}

//multiple return function

func fungsimultiple(x int, y int) (int, int) {
	var a = x * y
	var b = x / y
	return a, b
}
