package main

import "fmt"

func tipedata() {
	var numberuint8 uint8 = 255 //tipe data uint8 (0-255)
	fmt.Println(numberuint8)
	var numberuint16 uint16 = 65535 //tipe data uint16 (0 - 65535)
	fmt.Println(numberuint16)
	var numberuint32 uint32 = 4294967295 //tipe data uint32 (0 - 4294967295)
	fmt.Println(numberuint32)
	var numberuint64 uint64 = 18446744073709551615 //tipe data uint64 (0 - 18446744073709551615)
	fmt.Println(numberuint64)
	var numberint64 int64 = -9223372036854775807 // -9223372036854775808 - 9223372036854775807
	fmt.Println(numberint64)
	var numberint32 int32 = -2147483647 //-2147483648 - 2147483647
	fmt.Println(numberint32)
	var numberint16 int16 = -32768 //-32768 - 32767
	fmt.Println(numberint16)
	var numberint8 int8 = -128 //-128 - 127
	fmt.Println(numberint8)
	var desimal = 2.54 //float
	fmt.Println(desimal)
	var iya bool = true
	fmt.Println(iya)
	var pesan string = "Ini String yah"
	fmt.Println(pesan)
}
