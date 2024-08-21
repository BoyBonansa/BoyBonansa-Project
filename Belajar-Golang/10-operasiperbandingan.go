package main

import "fmt"

/*

	operasi perbandingan

	--> * operasi perbandingan adalah operasi untuk membandingkan dua buah data
	--> * operasi perbandingan adalah operasi yang menghasilkan nilai boolean ( benar atau salah )
	--> * jika hasil operasinya adalah benar, maka nilainya adalah true
	--> * jika hasil operasinya adalah salah, maka nilainya adalah false

	operator perbandingan

	operator -->	> 	= lebih dari
	operator -->	< 	= kurang dari
	operator -->	>=	= lebih dari sama dengan
	operator -->	<=	= kurang dari sama dengan
	operator --> 	== 	= sama dengan
	operator -->	!=	= tidak sama dengan

*/

func main() {

	// contoh kasus

	Nama1 := "Boy"     // membuat variable Nama1 dengan nilai "Boy"
	Nama2 := "Bonansa" // membuat variable Nama2 dengan nilai "Bonansa"

	var result bool = Nama1 == Nama2 // membuat variable result dengan nilai boolean Nama1 sama dengan Nama2
	fmt.Println(result)              // menampilkan result

	result = Nama1 != Nama2 // memberi nilai variable result dengan nilai boolean Nama1 tidak sama dengan Nama2
	fmt.Println(result)     // menampilkan result

	// contoh lain

	A := 10 // membuat variable A dengan nilai 10

	result = A < 9      // memberi nilai variable result dengan nilai boolean dari variable A kurang dari nilai 9
	fmt.Println(result) // menampilkan variable result

	result = A > 9      // memberi nilai variable result dengan nilai boolean dari variable A lebih dari nilai 9
	fmt.Println(result) // menampilkan variable result

	result = A == 9     // memberi nilai variable result dengan nilai boolean dari variable A sama dengan nilai 9
	fmt.Println(result) // menampilkan variable result

	result = A != 9     // memberi nilai variable result dengan nilai boolean dari variable A tidak sama dengan nilai 9
	fmt.Println(result) // menampilkan variable result
}
