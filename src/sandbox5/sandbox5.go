package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("-----OOP-----")
	c89()
	fmt.Println("-----method with nil receiver-----")
	c93()
	fmt.Println("-----embedded struct concept-----")
	c94()
}

type Point struct {
	X, Y float64
}

type Vector struct {
	X, Y float64
}

func (p *Point) Length() float64 {
	return math.Hypot(p.X, p.Y)
}

func (p *Point) MoveXTo(newX float64) {
	p.X = newX
}

func MoveYTo(p *Point, newY float64) {
	p.Y = newY
}

func c89() {
	p := Point{3, 4}
	fmt.Println(p)
	fmt.Println(p.Length())
	p.MoveXTo(20)
	fmt.Println(p.Length())
	MoveYTo(&p, 10)
	fmt.Println(p.Length())
	fmt.Println(p)

}

type BinaryTree struct {
	value       int
	left, right *BinaryTree
}

func (bt *BinaryTree) sum() int {
	if bt == nil {
		return 0
	}
	return bt.value + bt.left.sum() + bt.right.sum()
}

func c93() {
	bt := BinaryTree{
		value: 2,
		left:  &BinaryTree{value: 3, left: &BinaryTree{value: 1, left: nil, right: nil}, right: nil},
		right: &BinaryTree{value: 5, left: nil, right: nil},
	}

	fmt.Println(bt.sum()) //11
}

type Person struct {
	Name    string
	Surname string
}

func (p *Person) FullName() string {
	return p.Name + " " + p.Surname
}

type Employee struct {
	Person
	Id     string
	Office string
}

func (e *Employee) Detail() string {
	return "ID : " + e.Id + ". Office : " + e.Office + ". Fullname : " + e.FullName()
}

func (e *Employee) IsSameOffice(other *Employee) bool {
	return e.Office == other.Office
}

type Programmer struct {
	Employee
	Languages []string
}

func (p *Programmer) Detail() string {
	return "Programmer : " + p.Employee.Detail()
}

type Tester struct {
	Employee
	Tools []string
}

func (t *Tester) Detail() string {
	return "Tester : " + t.Employee.Detail()
}

func c94() {
	david := Person{
		Name:    "David",
		Surname: "Degea",
	}

	empDavid := Employee{
		Person: david,
		Id:     "123",
		Office: "Old Tafford",
	}

	progDavid := Programmer{
		Employee:  empDavid,
		Languages: []string{"golang", "Java", "C++"},
	}

	fmt.Printf("%+v\n", progDavid)

	ronaldo := Person{
		Name:    "Ronaldo",
		Surname: "Cristiano",
	}
	empRonaldo := Employee{
		Person: ronaldo,
		Id:     "456",
		Office: "Old Tafford",
	}

	testerRonaldo := Tester{
		Employee: empRonaldo,
		Tools:    []string{"robot", "Excel", "Note Pad"},
	}

	fmt.Printf("%+v\n", testerRonaldo)

	fmt.Println(empDavid.IsSameOffice(&empRonaldo))
	fmt.Println(progDavid.IsSameOffice(&(testerRonaldo.Employee)))
	fmt.Println(progDavid.FullName())
	fmt.Println(progDavid.Detail())

	// Method value
	davidDetail := progDavid.Detail
	fmt.Println("davidDetail : ", davidDetail())
	isSameOffice := (*Employee).IsSameOffice
	fmt.Println(isSameOffice(&empDavid, &empRonaldo))

}
