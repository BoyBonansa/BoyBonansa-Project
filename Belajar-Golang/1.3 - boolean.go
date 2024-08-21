package main

import "fmt"

/*  tipe data boolean adalah tipe data benar atau salah

--> * nilai boolean true	--> benar
--> * nilai boolean false 	--> salah

*/

func main() {
	fmt.Println("ini adalah tipe data boolean benar :", true)
	fmt.Println("ini adalah tipe data boolean salah :", false)

	// contoh kasus

	var Nilai1 bool     // membuat variable yang berisi tipe data boolean
	Nilai1 = true       // memberi nilai boolean true pada variable
	fmt.Println(Nilai1) // menampilkan variable

	var Nilai2 bool     // membuat variable yang berisi tipe data boolean
	Nilai2 = false      // memberi nilai boolean false pada variable
	fmt.Println(Nilai2) // menampilkan variable
}
