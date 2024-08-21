package main

import "fmt"

/*

	constant

	--> * constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
	--> * cara pembuatan constant mirip dengan variable, tetapi menggunakan syntac const
	--> * saat pembuatan constant, kita wajib langsung memberi nilai nya

*/

func main() {

	// contoh kasus

	const namaDepan string = "Boy"       // membuat constant dengan manual
	const namaBelakang = "Bonansa"       // membuat constant secara sederhana
	fmt.Println(namaDepan, namaBelakang) // menampilkan constan namaDepan dan namaBelakang

	// karena nilai constan itu adalah mutlak maka jika nilai nya diubah akan menjadi error

	// contoh kasus eror ( saya beri comment karena jika tidak akan terjadi eror)

	/*
		namaDepan = "Fatah"
		namaBelakang = "Rizki"
	*/

	// constant juka bisa di kasih multiple data seperti variable

	// contoh kasus
	// membuat multiple constant

	const (
		namaDepanLagi    = "adolf"
		namaBelakangLagi = "hitler"
	)

	fmt.Println(namaDepanLagi, namaBelakangLagi) // menampilkan constant namaDepanLagi dan namaBelakangLagi
}
