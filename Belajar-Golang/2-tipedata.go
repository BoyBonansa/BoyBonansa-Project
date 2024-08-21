package main

import "fmt"

/** tipe data number
1. tipe data integer ( mulai dari negatif )

	a. int8 	--> nilai minimum -128
					nilai maximum  127

	b. int16 	--> nilai minimum -32768
					nilai maximum  32767

	c. int32	--> nilai minimum -2147483648
				    nilai maximum  2147483647

	d. int64	--> nilai minimum -9223372036854775808
					nilai maximum  9223372036854775807

2. tipe data integer ( mulai dari positif )

	a. uint8	--> nilai minimum 0
					nilai maximum 255

	b. uint16	--> nilai minimum 0
					nilai maximum 65535

	c. uint32 	--> nilai minimum 0
					nilai maximum 4294967295

	d. uint64	--> nilai minimum 0
					nilai maximum 18446744073709551615

3. tipe data floating point

	a. float32	--> nilai minimum 1.18 x 10^38
					nilai maximum 3.4 x 10^38

	b. float64	--> nilai minimum 2.23 x 10^308
					nilai maximum 1.80 x 10^308

	c. complex64	--> complex numbers with float32 real and imaginary parts

	d. complex128	--> complex numbers with float64 real and imaginary parts

4. alias ( nama lain dari tipe data )

	a. byte		--> uint8

	b. rune		--> int32

	c. int		--> minimal int32

	d. uint		--> minimal  uint32
*/

func main() {
	fmt.Println("ini adalah tipe data int :", 7)
	fmt.Println("ini adalah tipe data float :", 10.5)
}
