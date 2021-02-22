package main

import "fmt"

func tugas4() {
	var datatb = map[string]int{"Andi": 167, "Budi": 171}
	tampilkandata(datatb)
}

func tampilkandata(datatb map[string]int) {
	for key, value := range datatb {
		fmt.Println(key, ":", value, " cm")
	}
}
