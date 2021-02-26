package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func tugas7() {
	runtime.GOMAXPROCS(2)
	var angka = 391
	var kata = "Huruf"
	var reflectAngka = reflect.ValueOf(angka)
	var reflectKata = reflect.ValueOf(kata)
	go bacareflect1(reflectAngka.Type().Name())
	bacareflect2(reflectKata.Type().Name())

	//Blocking untuk menahan process
	var input string
	fmt.Scanln(&input)
}

func bacareflect1(x string) {
	fmt.Println(x)
}
func bacareflect2(x string) {
	fmt.Println(x)
}
