package main

import "fmt"

/*

	operasi boolean

	operator -->	&& = dan
	operator -->	|| = atau
	operator -->	!  = kebalikan

	operasi && ( dan )

	syarat dari operasi && adalah jika salah satu nilai false maka hasilnya adalah false

	true && true 	= true
	true && false 	= false
	false && true	= false
	false && false 	= false

	operasi || ( atau )

	syarat dari operasi || adalah jika salah satu nilai true maka hasilnya adalah true

	true || true	= true
	true || false	= true
	false || true	= true
	false || false 	= false

	operasi ! ( kebalikan )

	!true = false
	!false = true

*/

func main() {

	// contoh

	nilaiMateri := 90 // membuat variable nilaiMateri dengan nilai 90
	nilaiSikap := 95  // membuat variable nilaiSikap dengan nilai 95

	kkmMateri := nilaiMateri > 75 // membuat variable kkmMateri dengan nilai nilaiMateri lebih dari 75
	kkmSikap := nilaiSikap > 75   // membuat variable kkmSikap dengan nilai nilaiSikap lebih dari 75

	lulus := kkmMateri && kkmSikap // membuat variable lulus dengan nilai operasi boolean kkmMateri dan kkmSikap

	fmt.Println(lulus) // menampilkan variable lulus

	nilaiSikap = 20                // memberi nilai 20 pada variable nilaiSikap
	kkmSikap2 := nilaiSikap > 75   // membuat variable nilaiSikap2 dengan nilaiSikap lebih dari 75
	lulus = kkmMateri && kkmSikap2 // membuat variable lulus dengan nilai operasi boolean kkmMateri dan kkmSikap2

	fmt.Println(lulus) // menampilkan variable lulus
}
