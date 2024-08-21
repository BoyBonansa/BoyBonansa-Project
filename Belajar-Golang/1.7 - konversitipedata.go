package main

import "fmt"

/*

	konversi tipe data

	--> * di go-lang terkadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
	--> * seperti mengkonversi tipe data int32 ke int64, float32 ke float64, dan sebagainya

*/

func main() {

	// contoh kasus

	var A int32 = 32768    // membuat variable A dengan tipe data int32 nilai nya 32768
	var B int64 = int64(A) // membuat variable B dengan nilai kita ambil dari variable A lalu dikonversi menjadi int64
	// int32 --> int64 ini bisa karena mengkonversi ke jangkauan yang lebih luas

	var C int16 = int16(A) // membuat variable C dengan nilai kita ambil dari variable A lalu dikonversi menjadi int16
	// int32 --> int16 ini akan merubah nilainya karena mengkonversi ke jangkauan yang lebih kecil
	fmt.Println(A) // menampilkan variable A
	fmt.Println(B) // menampilkan variable B
	fmt.Println(C) // menampilkan variable C

	// contoh kasus lain

	var Nama = "Boy Bonansa" // membuat variable Nama dengan nilai "Boy Bonansa"
	var D = Nama[2]          // membuat variable D dengan nilai diambil dari nama lalu yang diambil karakter ke-2
	fmt.Println(Nama)        // menampilkan variable nama
	fmt.Println(D)           // menampilkan variable D

	var E = string(D) // membuat variable E dengan nilai diambil dari variable D lalu dikonversi tipe datanya menjadi string
	fmt.Println(E)    // menampilkan variable E
}
