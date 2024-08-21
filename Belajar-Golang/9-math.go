package main

import "fmt"

/*

	operasi matematika

	operator --> +	= penjumlahan
	operator --> -	= pengurangan
	operator --> *	= perkalian
	operator --> /	= pembagian
	operator --> %	= sisa pembagian

*/

func main() {

	// contoh

	var a int = 10 // membuat variable dengan nama a tipe data int dan nilai nya 10
	var b int = 10 // membuat variable dengan nama b tipe data int dan nilai nya 10

	var c = a + b  // membuat variable dengan nama c dan nilainnya variable a ditambah variable b
	fmt.Println(c) // menampilkan variable c

	c = a - b      // memberi nilai variable c dengan nilai variable a dikurang variable b
	fmt.Println(c) // menampilkan variable c

	c = a * b      // memberi nilai variable c dengan nilai variable a dikali variable b
	fmt.Println(c) // menampilkan variable c

	c = a / b      // memberi nilai variable c dengan nilai variable a di bagi variable b
	fmt.Println(c) // menampilkan variable c

	// contoh kasus lain

	var d int = 5 // membuat variable dengan nama d tipe data int dan nilai nya 5

	c = a + b - 1*d // memberi nilai variable c dengan nilai dari variable a ditambah variable b dikurang 19 dikali variable d
	fmt.Println(c)  // menampilkan variable c

	/*

		augmented assignments

		--> * augmented assignments adalah melakukan operasi ke diri sendiri atau ke variable itu sendiri

		operasi matematika -->	a = a + 10		augmented assignments -->	a += 10
		operasi matematika -->	a = a - 10		augmented assignments -->	a -= 10
		operasi matematika --> 	a = a * 10		augmented assignments --> 	a *= 10
		operasi matematika --> 	a = a / 10		augmented assignments --> 	a /= 10
		operasi matematika --> 	a = a % 10		augmented assignments -->	a &= 10

	*/

	// contoh kasus

	var aa int = 100 // membuat variable dengan nama aa tipe data int dan nilai nya 100
	aa += 10         // membuat augmented assignments untuk variable itu sendiri

	fmt.Println(aa) // menampilkan variable

	/*

		unary operator

		--> * shortcut untuk membuat angka bertambah 1 atau sebaliknya

		operator -->	++		keterangan --> a = a + 1 ( ditambah satu )
		operator -->	--		keterangan --> a = a - 1 ( dikurang satu )
		operator -->	-		keterangan --> negative
		operator --> 	+		keterangan --> positive
		operator --> 	! 		keterangan --> boolean kebalikan

	*/

	// contoh kasus

	j := 1 // membuat variable dengan nilai 1
	j++    // memberi operasi unary dengan menambah 1 totalnya 2

	fmt.Println(j) // menampilkan variable j

	j++ // memberi operasi unary dengan menambah 1 totalnya 3
	j++ // memberi operasi unary dengan menambah 1 totalnya 4

	fmt.Println(j) // menampilkan variable j
}
