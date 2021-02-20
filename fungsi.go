package main

import "fmt"

func fungsi() {
	fmt.Println(fungsi_pertama(10, 5)) //Pemanggilan fungsi
}

func fungsi_pertama(x int, y int) int { //ini inisiasi fungsi mengembalikan nilai string
	var z = x + y

	return z
}
