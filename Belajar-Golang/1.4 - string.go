package main

import "fmt"

/*
	 	tipe data string

		-->	* string adalah tipe data kumpulan karakter
		-->	* jumlah karakter di dalam string bisa nol sampai tidak terhingga
		-->	* tipe data string di go-lang direpresentasikan ( syntax ) dengan kata kunci string
		-->	* nilai data string di go-lang selalu diawali dan diakhiri dengan karakter --> " <-- ( petik dua)

		function untuk string

		1. function --> len("string")		keterangan --> menghitung jumlah karakter di string

		2. function --> "string"[number]	keterangan --> mengambil karakter pada posisi yang ditentukan
*/
func main() {

	// contoh kasus string

	fmt.Println("Boy Bonansa") // menampilkan data string "Boy Bonansa"

	// contoh kasus function untuk string

	fmt.Println(len("Boy Bonansa")) // menghitung total karakter pada data string "Boy Bonansa"
	fmt.Println("Boy Bonansa"[3])   // mengambil karakter pada posisi [3] dari data string "Boy Bonansa" ditampilkan berupa byte

	// contoh kasus lain

	var Nama string                // membuat variable dengan tipe data string
	Nama = "Verlita Eka Febrianti" // memberi value pada variable

	fmt.Println(len(Nama), Nama[4]) // menampilkan len dari variable dan juga mengambil byte variable [4]
}
