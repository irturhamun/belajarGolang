package main

import "fmt"

func fungsivariadic() {
	var rataan = hitung(2, 35, 6, 3, 1, 231, 53, 2, 121, 1)
	var pesan = fmt.Sprintf("Rataan : %.2f", rataan)
	fmt.Println(pesan)
}

func hitung(angka ...int) float64 {
	var total int = 0
	for _, angka := range angka {
		total += angka
		fmt.Println(angka)
	}
	var avg = float64(total) / float64(len(angka))
	return avg
}
