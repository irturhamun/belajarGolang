package main

import (
	"fmt"
	"reflect"
)

func reflectGo() {
	var number = 10
	var reflectNumber = reflect.ValueOf(number) //bisa melihat tipe data dan lainnya
	fmt.Println("Tipe Variable :", reflectNumber.Type())

	if reflectNumber.Kind() == reflect.Int {
		fmt.Println("Nilai Variable :", reflectNumber.Int())
	}
}
