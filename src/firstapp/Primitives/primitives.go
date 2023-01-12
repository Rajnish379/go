// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n",n, n)
	x := 1 == 1
	y := 2 == 3
	var m uint16 = 45
	var k bool 
	// Default value of an uninitialized variable is false
	fmt.Printf("%v, %T",k, k)
	fmt.Printf("%v,%T",m,m)
	fmt.Printf("%v, %T\n", x,x)
	fmt.Printf("%v, %T\n",y, y)
	a := 10 //1010
	b := 3 // 0011
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a % b)
	fmt.Println(a&b) // 0010
	fmt.Println(a|b) // 1011
	fmt.Println(a^b) // 1001
	// a & ~b
	fmt.Println(a &^ b) //1000
	var c int = 10
	var d int8 = 3
	// Cannot add two variables of different numerical type
	// fmt.Println(c+d)
	// Have to type convert the variable to make it work
	fmt.Println(c+int(d))
	e := 8  // 2^3
	fmt.Println(e << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(e >> 3) // 2^3 // 2^3 = 2^0
	p := 3.14
	p = 13.7e72
	p = 2.1E14
	fmt.Printf("%v, %T\n",p, p)
	f := 10.2
	g := 3.7
	// Remainder, bitwise, and bitshifting operators are not available on floating type numbers
	// If u need them u have to work on integer type numbers only
	fmt.Println(f / g)
	fmt.Println(f * g)
	fmt.Println(f + g)
	fmt.Println(f - g)
	var v complex64 = 1 + 2i
	var u complex128 = complex(5,12)
	fmt.Printf("%v, %T\n", v, v)
	fmt.Printf("%v, %T\n", u, u)
	// If v is complex64 real and imag parts will be float32 each
	// v is complex128 real and imag parts will be float64 each
	fmt.Printf("%v, %T\n", real(v), real(v))
	fmt.Printf("%v, %T\n", imag(v), imag(v))
	q := 1+2i
	w := 2+5.2i
	fmt.Println(q+w)
	fmt.Println(q-w)
	fmt.Println(q*w)
	fmt.Println(q/w)
	s := "this is a string"
	s2 := "this is also a string"
	byte := []byte(s)
	fmt.Printf("%v, %T\n", s,s)
	// Strings are aliases of bytes so we have to convert them back to string using string type converter to return the character at the specific index
	fmt.Printf("%v, %T\n", (s[2]),s[2])
	fmt.Printf("%v, %T\n", string(s[2]),s[2])
	// Cannot assign string to byte for changing value at that particular position of string
	// s[2] = "u"
	fmt.Printf("%v, %T\n", s+s2,s+s2)
	fmt.Printf("%v, %T\n", byte,byte)
	//Runes are assigned with single quotes while strings are assigned double quotes
	r := 'a'
	var z rune = 'b'
	fmt.Printf("%v,%T\n",r,r)
	fmt.Printf("%v,%T\n",z,z)
}
