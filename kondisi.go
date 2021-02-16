package main

import "fmt"

func kondisi() {
	var nilai = 7
	//Seleksi Kondisi menggunakan IF/ELSE Statement
	if nilai == 10 {
		fmt.Println("Perfect")
	} else if nilai >= 7 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	//Seleksi expressionsi kondisi menggunakan Switch Case
	var PassGrade = 6
	switch PassGrade {
	case 7:
		fmt.Println("A")
	case 6:
		fmt.Println("B")
	case 5:
		fmt.Println("C")
	default:
		fmt.Println("Ungraded")

	}
}
