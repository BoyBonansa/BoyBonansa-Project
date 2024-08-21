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

	/*

		kata kunci var

		--> * di go-lang kata kunci var saat membuat variable tidak wajib
		--> * asalkan saat membuat variable kita langsung memberi nilai datanya
		--> * agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci :=
		--> * sebenarnya ini berfungsi hanya untuk menyederhanakan penulisan kode saja

	*/

	// contoh kasus

	Sahabat := "ahmad rizki simbolon" // membuat variable tanpa menggunakan kata kunci var dan langsung di beri nilai
	fmt.Println(Sahabat)              // menampilkan variable

	Sahabat = "ahmad faiz al-fatah" // mengganti nilai variable sahabat, tetapi hanya menggunakan = tidak boleh menggunakan := karena := hanya untuk deklarasi awal
	fmt.Println(Sahabat)            // menampilkan variable

	/*

		multiple variable

		--> * di go-lang kita bisa membuat variable secara sekaligus banyak
		--> * code yang dibuat akan lebih bagus dan mudah dibaca

	*/

	// contoh kasus

	// membuat multiple variable Nama1 dan Nama2 lalu langsung diberi nilai nya

	var (
		Nama1 = "adolf hitler"
		Nama2 = "joseph stalin"
	)

	fmt.Println(Nama1) // menampilkan variable Nama1
	fmt.Println(Nama2) // menampilkan variable Nama2
}
