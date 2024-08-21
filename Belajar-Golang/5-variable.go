package main

import "fmt"

/*

	variable

	--> * variable adalah tempat untuk menyimpan data
	--> * variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
	--> * di go-lang variable hanya bisa menyimpan tipe data yang sama
	--> * untuk membuat variable kita menggunakan syntax var, lalu diikuti dengan nama variable dan tipe datanya

*/

func main() {

	// contoh kasus

	var Nama string      // membuat variable dengan nama Nama dan tipe datanya string
	Nama = "Boy Bonansa" // memberi nilai "Boy Bonansa" pada variable Nama
	fmt.Println(Nama)    // menampilkan variable Nama

	var Umur int      // membuat variable dengan nama Umur dan tipe datanya int
	Umur = 20         // memberi nilai 20 pada variable Umur
	fmt.Println(Umur) // menampilkan variable Umur

	var Status bool                               // membuat variable dengan nama Status dan tipe datanya boolean
	Status = true                                 // memberi nilai true pada variable Status
	fmt.Println("Sudah memiliki pacar :", Status) // menampilkan variable Status dan ditambahkan string

	/*

		tipe data variable

		--> * saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
		--> * namun jika kita langsung memberi nilai pada variable, maka kita tidak wajib menyebutkan tipe datanya

	*/

	// contoh

	var Pacar = "verlita eka febrianti" // membuat variable dengan nama pacar tetapi langsung memberi nilai string "verlita eka febrianti" pada variable
	fmt.Println(Pacar)                  // menampilkan variable
}
