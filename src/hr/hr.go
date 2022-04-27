package hr

type Person struct {
	Name string
	Age  int
}

// anonymous field
type Employee struct {
	Person
	Designation string
}
