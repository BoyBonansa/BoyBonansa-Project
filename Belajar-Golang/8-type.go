package main

import "fmt"

/*

	type declarations

	--> * type declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
	--> * type declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada
	--> * tujuannya adalah agar lebih mudah dimengerti
	--> * kesimpulannya adalah kita membuat nama alias dari tipe data yang kita inginkan

*/

func main() {

	// contoh kasus

	type NoHP string // membuat type dengan nama NoHP sebagai alias dari tipe data string

	var boy NoHP = "0813 9341 5408"   // membuat variable dengan nama boy tipe data NoHP dan memberi nilai nya
	fmt.Println(boy)                  // menampilkan variable boy
	fmt.Println(NoHP("081308130813")) //menampilkan tipe data NoHp dan memberi nilainya
}
