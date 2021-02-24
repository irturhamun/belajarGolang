package library

import (
	"fmt"
	"math/rand"
)

func iniPrivat() {
	//kalau fungsinya diawali huruf kecil jadinya Privat
	fmt.Println("Ini fungsi Private")
}

/* IniPublic merupakan fungsi untuk jadi Public*/
func IniPublic() {
	//Kalau fungsinya diawali huruf besar jadinya Public
	fmt.Println("Ini fungsi Public")
	iniPrivat()
}

type pelajar2 struct {
	nama string
	umur int
}

func dataMahasiswa() {
	var n1, n2, n3 pelajar2
	n1.nama = "Ahmad"
	n1.umur = rand.Intn(100)
	n2.nama = "Abi"
	n2.umur = rand.Intn(100)
	n3.nama = "Zaka"
	n3.umur = rand.Intn(100)
	fmt.Println(n1.nama, n1.umur)
	fmt.Println(n2.nama, n2.umur)
	fmt.Println(n3.nama, n3.umur)
}

func TampilkanData() {
	dataMahasiswa()
}
