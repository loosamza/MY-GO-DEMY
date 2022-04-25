package main

import (
	"fmt"
	"math"
	"math/rand"
	"prime"
	"weight"

	weightGH "github.com/loosamza/wieght"
)

// Section 4

func main() {
	fmt.Println("----------")
	c22()
	fmt.Println("----------คอนเซปเรื่อง pointer----------")
	c24()
	fmt.Println("----------function return pointer----------")
	c25()
	fmt.Println("----------function receive pointer และ คอนเซป alias----------")
	c26()
	fmt.Println("----------การประกาศตัวแปรแบบไม่มีชื่อ (new function)----------")
	c27()
	fmt.Println("----------type declaration----------")
	c28()
	fmt.Println("----------package and import----------")
	c33()
	fmt.Println("----------package and initialization function----------")
	c35()

}

func c22() {
	x := float64(1)
	b := math.Sqrt(2)
	c := x + b
	fmt.Println(c)
}

func c24() {
	x := 70
	var p *int
	p = &x
	fmt.Println("&x", &x)
	fmt.Println("*(&x)", *(&x))
	fmt.Println("p", p)
	fmt.Println("&p", &p)
	fmt.Println("*p", *p)
}

func fp() *int {
	x := 4
	return &x
}

func c25() {
	x := fp()
	// false forever because fp alway return new address
	fmt.Println(x == fp())

}

func c26() {
	x := 2
	// double(x) cannot use x (type int) as type *int in argument to double
	double(&x)
	fmt.Println(x)
}

func double(alias *int) {
	*alias = *alias * 2
}

func newIntPointer() *int {
	var x int
	return &x
}
func c27() {
	x := newIntPointer()
	fmt.Println(&x)
	fmt.Println(newIntPointer() == newIntPointer()) // false
	// ทำงานเหมือนกัน
	fmt.Println(new(string) == new(string)) // false
	fmt.Println(new(int) == new(int))       // false
}

type KG float64
type LB float64

func (lb LB) toKG() KG {
	return KG(lb / 2.2046226218)
}

func (kg KG) toLB() LB {
	return LB(kg / 0.453592)
}

func (kg KG) toString() string {
	return fmt.Sprint(kg)
}

func (lb LB) toString() string {
	return fmt.Sprint(lb)
}

func c28() {
	k := KG(3)
	// implicit
	result := k + 3.0
	fmt.Println(result)
	fmt.Printf("%T\n", result)

	b := LB(3)
	// CAST
	fmt.Println(k == KG(b))
	// สร้างคุณลักษณะเพิ่มเติม
	fmt.Println(k > b.toKG())
	fmt.Println(LB(3).toKG().toLB())
	fmt.Println(LB(3) == LB(3).toKG().toLB())
	fmt.Println(k.toString())
	fmt.Println(b.toString())
}

func c33() {
	k := weightGH.KG(2)
	k2 := weight.KG(2)

	fmt.Println(weightGH.KgToLbCal(k))

	fmt.Println(weight.KgToLbCal(k2))

	//fmt.Println("KG from different source", k == k2) invalid operation: k == k2 (mismatched types "github.com/loosamza/wieght".KG and "weight".KG)

}

func c35() {
	for i := 0; i < 100; i++ {
		x := rand.Intn(1000000)
		fmt.Printf("%d , %t\n", x, prime.IsPrime(x))
	}
	// fmt.Printf("%d , %t\n", 13, prime.IsPrime(5))
}
