package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Section 5
func main() {
	fmt.Println("-----string representation-----")
	c45()
	fmt.Println("-----การแปลง string ระหว่าง unicode กับ utf-8-----")
	c46()
	fmt.Println("-----string conversion-----")
	c47()
	fmt.Println("-----range-loop และ pacakge สำหรับ string-----")
	c49()
	fmt.Println("----constant-----")
	c50()
}

func c45() {
	x := "hi-สวัสดี"
	y := []byte{0x68, 0x69, 0x2d, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	fmt.Println(x, len(x))
	fmt.Println(string(y))
	fmt.Println(string(y[0]))                 // h
	fmt.Println(string(y[1]))                 // i
	fmt.Println(string(y[2]))                 // -
	fmt.Println(string(y[3]), string(y[3:6])) // ส
	fmt.Println(string(y[4]), string(y[6:9])) // ว

	/*
		Real output
		h
		i
		-
		à

		why ?????
	*/

}

func c46() {

	x := "hi-สวัสดี"
	y := []byte{0x68, 0x69, 0x2d, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	z := []rune(x)
	fmt.Println(string(0xE2A))
	fmt.Println("utf-8", "\xe0\xb8\xaa")
	fmt.Println(len(x))
	fmt.Println(len(y))
	fmt.Println(len(z))
	fmt.Printf("%q\n", z[4])

}

func c47() {

	// Case 1 : Integer to String
	ex1 := string(0x265e)
	fmt.Println("Case 1: Integer -> String")
	for i := 0; i < len(ex1); i++ {
		fmt.Println("%x ", ex1[i])
	}

	fmt.Println("\n\xe2\x99\x9e")
	fmt.Println("\xe2\x99\x9e" == ex1)

	// Case 2 : slice of byte to string
	fmt.Println("\nCase 2: []byte -> string")
	ex2 := []byte{0xe2, 0x99, 0x9e}
	ex2String := string(ex2)
	fmt.Println(ex2String)

	// Case 3 : slice of rune to string
	fmt.Println("\nCase 3: []rune -> string")
	ex3 := []rune{0x2654, 0x2655, 0x2656, 0x2657, 0x2658, 0x2659}
	fmt.Println(string(ex3))

	// Case 4 : slice of byte to string
	fmt.Println("\nCase 4: string -> []byte")
	ex4 := []byte("Hello♕")
	fmt.Println(ex4)

	// Case 5 : string to a slice of runes
	fmt.Println("\nCase 5 : string -> []rune")
	ex5 := []rune("Hello♕")
	fmt.Println(ex5)

}

func c49() {
	x := "ทดสอบ"
	for i, v := range x {
		fmt.Printf("%d, %c\n", i, v)
		fmt.Println(i, ",", string(v))

		fmt.Println("utf8.RuneCountInString(x)", utf8.RuneCountInString(x))
		fmt.Println(utf8.DecodeRuneInString(x))

		for i := 0; i < len(x); {
			r, s := utf8.DecodeRuneInString(x[i:])
			i += s
			fmt.Printf("%c-", r)
		}
	}

	finder := "สอ"
	fmt.Println()
	fmt.Println("bytes.Count([]byte(x), []byte(finder))", bytes.Count([]byte(x), []byte(finder)))
	fmt.Println("strings.Count(x, finder)", strings.Count(x, finder))

	buff := bytes.Buffer{}
	buff.WriteRune('y')
	buff.WriteString("o")
	fmt.Println(buff.String())

	buff2 := strings.Builder{}
	buff2.WriteRune('y')
	buff2.WriteString("o")
	fmt.Println(buff2.String())

	atoi, _ := strconv.Atoi("123")
	itoa := strconv.Itoa(123)
	fmt.Println(atoi, reflect.TypeOf(atoi))
	fmt.Println(itoa, reflect.TypeOf(itoa))

	fmt.Println(strconv.ParseBool("tRue"))
	fmt.Println(strconv.ParseBool("5"))

	fmt.Println("unicode.IsDigit('๓')", unicode.IsDigit('๓'))
	fmt.Println("unicode.IsUpper('A')", unicode.IsUpper('A'))
	fmt.Println("unicode.In('ด', unicode.Thai)", unicode.In('ด', unicode.Thai))
	fmt.Println("unicode.IsDigit('a')", unicode.IsDigit('a'))

}

func c50() {
	const persons = (4)
	//const persons = int(4) // 2.01 (untyped float constant) truncated to intcompiler
	toffee := 5 / persons
	cost := 2.01 / persons

	fmt.Println(toffee, reflect.TypeOf(toffee))
	fmt.Println(cost, reflect.TypeOf(cost))

	// const All = 1
	// const Free = 2
	// const Discount = 3
	// const (
	// 	All = 1
	// 	Free = 2
	// 	Discount = 3
	// )
	// const (
	// 	All = iota + 1
	// 	Free
	// 	Discount
	// )
	// fmt.Println(All, Free, Discount) // 1 2 3
	// const (
	// 	All = iota + 1
	// 	Free
	// 	Discount = 99
	// )
	// fmt.Println(All, Free, Discount) // 1 2 99
	const (
		All      = iota + 1
		Free     = 99
		Discount = iota
	)
	fmt.Println(All, Free, Discount) // 1 99 2
}
