package main

import "fmt"

/*

	array

	--> * array adalah tipe data yang berisikan kumpulan data dengan tipe data yang sama
	--> * saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh array
	--> * daya tampung array tidak bisa bertambah setelah array dibuat

	saat kita membuat tipe data array bayangkan bahwa array itu adalah lemari yang terdapat rak-rak nya

	index di array dimulai dari 0, jika jumlah array adalah [3] maka indexnya 0,1,2 karena dimulai dari 0

	index -->	[0]	= "fatah"
	index -->	[1] = "boy"
	index -->	[2] = "vigo"

*/

func main() {

	// contoh kasus

	var Nama [3]string // membuat variable array yang bernama Nama dengan jumlah data 3 dan tipe data string
	Nama[0] = "fatah"  // mengisi index [0] dengan nilai "fatah"
	Nama[1] = "boy"    // mengisi index [1] dengan nilai "boy"
	Nama[2] = "vigo"   // mengisi index [2] dengan nilai "vigo"

	fmt.Println(Nama)    // menampilkan data array Nama
	fmt.Println(Nama[1]) // menampilkan data array Nama index[1]

	Nama[2] = "rizki" // mengubah nilai di index[2] pada array Nama dengan nilai "rizki"
	fmt.Println(Nama) // menampilkan data array Nama ( karena tadi merubah nama[2] maka data di dalam array juga ikut berubah )

	// membuat array secara langsung
	// contoh

	Nilai := [5]int{ // membuat array secara langsung lebih menyederhanakan pengkodean
		10, // memberi nilai 10 pada index[0]
		20, // memberi nilai 20 pada index[1]
		30, // memberi nilai 30 pada index[2]
		40, // memberi nilai 40 pada index[3]
		50, // memberi nilai 50 pada index[4]
	}

	fmt.Println(Nilai) // menampilkan data array Nilai

	// membuat array yang lebih sederhana lagi
	// contoh

	Angka := [3]int{5, 10, 15} // membuat array dengan jumlah data 3 dan tipe datanya int
	fmt.Println(Angka)         // menampilkan data array
}
