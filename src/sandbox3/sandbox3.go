package main

import (
	"encoding/json"
	"fmt"
	"hr"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// Section 6 Data Structure

func main() {

	fmt.Println("-----arrays-----")
	c53()
	fmt.Println("-----slice and concept-----")
	c55()
	fmt.Println("-----map-----")
	c58()
	fmt.Println("-----struct-----")
	c59()
	fmt.Println("-----struct embedding and anonymous field-----")
	c61()
	fmt.Println("-----struct and recursive-----")
	c62()
	fmt.Println("-----JSON and Unmarshal-----")
	c63()
	fmt.Println("-----JSON with httpGET and json encode,decode-----")
	c65()
	fmt.Println("-----text template solution actions and more-----")
	c68()
	fmt.Println("-----HTML template-----")
	c69()
}

func c53() {
	fruits := [...]string{
		"Banana",
		"Mango",
		"Orange",
		"Papaya",
	}

	// fruits := [...]string{
	// 	"Banana",
	// 	"Mango",
	// 	"Orange",
	// 	"Papaya",
	// }

	// fruits := []string{
	// 	4: "Banana",
	// 	1: "Mango",
	// 	2: "Orange",
	// 	3: "Papaya",
	// } len 4

	fmt.Println(len(fruits), fruits, fruits[2], fruits[3])

	twoSlots := [2]int{}
	threeSlots := [3]int{}
	fmt.Println("Two slots\t", reflect.TypeOf(twoSlots))
	fmt.Println("Three slots\t", reflect.TypeOf(threeSlots))
	//fmt.Println(twoSlots == threeSlots) cannot compare twoSlots == threeSlots (mismatched types [2]int and [3]int)

	fmt.Println("expect true", [2]int{1, 2} == [2]int{1, 2})
	fmt.Println("expect false", [2]int{1, 2} == [2]int{2, 1})

	dubFruits := fruits
	fmt.Println("fruits", fruits)
	fmt.Println("dubFruits", dubFruits)
	fmt.Println("dubFruits[0] = \"Watermelon\"")

	fmt.Println("fruits", fruits)
	fmt.Println("dubFruits", dubFruits)

	ptrFruits := &fruits
	fmt.Println("ptrFruits", *ptrFruits)
	fmt.Println("fruits", fruits)
	fmt.Println("ptrFruits[0] = \"Watermelon\"")
	ptrFruits[0] = "Watermelon"
	fmt.Println("ptrFruits", *ptrFruits)
	fmt.Println("fruits", fruits)

	// muti dimention array declare
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(a)

}

func c55() {
	fruits := [5]string{"apple", "banana", "papaya", "orange", "mango"}
	myFavor := fruits[1:4]
	fmt.Println(len(myFavor), cap(myFavor), myFavor)

	yourFavor := myFavor
	fmt.Println(len(yourFavor), cap(yourFavor), yourFavor)
	yourFavor[0] = "guava"
	fmt.Println(len(yourFavor), cap(yourFavor), yourFavor)
	fmt.Println(len(myFavor), cap(myFavor), myFavor)

	x := [10]int{}
	a := x[0:5]
	b := x[5:7]
	for i := range a {
		a[i] = i * i
	}
	b[0] = 98
	b[1] = 99
	a = append(a, b...)
	//a = append(a, b[0], b[1]) same
	a[len(a)-1] = 25
	a = append(a, 71, 72, 73, 74)
	fmt.Println("x :\t", x)
	fmt.Println("a :\t", len(a), cap(a), a)
	fmt.Println("b :\t", len(b), cap(b), b)

	slice := append([]byte("hello "), "world"...)
	fmt.Println("slice :\t", len(slice), cap(slice), slice)

}

func c58() {

	// Create new map
	// literal styles
	items := map[string]int{
		"pencil": 5,
		"pen":    10,
	}
	fmt.Println(items)
	fmt.Println(items["pen"])
	delete(items, "pen")
	fmt.Println(items)
	x := items
	x["ruler"] = 7
	fmt.Println(x)
	fmt.Println(items)

	// make function styles
	orders := make(map[string]int)
	orders["pen"] = 15
	orders["pencil"] = 16
	fmt.Println(orders)
	fmt.Println(orders["pencil"])
	delete(orders, "pencil")
	fmt.Println(orders)

	v, ok := orders["xxx"]
	if !ok {
		fmt.Println("not exist", v)
	} else {
		fmt.Println("Exist", v)
	}
	//address := &items["ruler"]  cannot take address of items["ruler"]
	for k, v := range items {
		fmt.Println(k, v)
	}

	fmt.Println("len(items)", len(items))

}

func c59() {
	// Toy
	x := struct {
		name string
		Age  int
	}{Age: 28, name: "Toy Godemy"}

	fmt.Println(x)
	fmt.Println(x.name)
	fmt.Println(x.Age)

	// Real

	y := hr.Person{Age: 28, Name: "Real Godemy"}

	fmt.Println(y)
	fmt.Println(y.Name)
	fmt.Println(y.Age)

	z := y
	z.Name = "New Godemy"
	fmt.Println(y)
	fmt.Println(z)

	a := &y
	a.Name = "New Godemy"
	fmt.Println(y)
	fmt.Println(a)

	c1 := hr.Person{Age: 28, Name: "Real Godemy"}
	c2 := hr.Person{Age: 28, Name: "Real Godemy"}

	n := new(hr.Person)
	n.Name = "Godemy"
	n.Age = 22
	fmt.Println("n1", *n)

	n = &hr.Person{Name: "R", Age: 22}
	fmt.Println("n2", &n)

	fmt.Println("Compare struct c1&c2", c1 == c2)

}

func c61() {

	filicity := hr.Employee{
		Person:      hr.Person{Name: "Filicity", Age: 22},
		Designation: "Programmer",
	}
	fmt.Printf("%+v\n", filicity)
	fmt.Println(filicity.Name)
	fmt.Println(filicity.Person.Name)
	fmt.Println(filicity.Age)

}

//invalid recursive type BinaryTree
// type BinaryTree struct {
// 	value int
// 	left  BinaryTree
// 	right BinaryTree
// }

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func c62() {

	root := BinaryTree{value: 2}
	left := BinaryTree{value: 1}
	right := BinaryTree{value: 3}

	root.left = &left
	root.right = &right

	showDF(&root)

}

func showDF(node *BinaryTree) {
	if node != nil {
		showDF(node.left)
		fmt.Println(node.value)
		showDF(node.right)
	}
}

var data = `[{
    "userId": 1,
    "id": 1,
    "title": "delectus aut autem",
    "completed": false
  },
  {
    "userId": 1,
    "id": 2,
    "title": "quis ut nam facilis et officia qui",
    "completed": false
  },
  {
    "userId": 2,
    "id": 5,
    "completed": false
  } 
  ]`

type Todo []struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed *bool  `json:"completed,omitempty"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type User []struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

func c63() {

	dataStruct := Todo{}
	v := &dataStruct
	json.Unmarshal([]byte(data), v)
	fmt.Println(dataStruct)
	fmt.Println(len(dataStruct))
	// dataStruct[0].Completed = true
	dataStruct[1].Completed = &[]bool{true}[0]

	result, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return
	}
	fmt.Println(string(result))

}

func c65() {
	// todos
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	dataStruct := []Todo{}
	v := &dataStruct
	json.Unmarshal(data, v)
	fmt.Println(len(dataStruct))

	result, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return
	}
	fmt.Println(string(result))

	// users with json decord,encode

	resp2, err2 := http.Get("https://jsonplaceholder.typicode.com/users")
	if err2 != nil {
		return
	}
	jsonDecoder := json.NewDecoder(resp2.Body)
	dataStruct2 := User{}

	jsonDecoder.Decode(&dataStruct2)

	resp2.Body.Close()
	fmt.Println(len(dataStruct2))
	dataStruct2[0].Name = "Scontz Ja"
	jsonEncoder := json.NewEncoder(os.Stdout)
	jsonEncoder.Encode(&dataStruct2)

}

func upperCase(s string) string {
	return strings.ToUpper(s)
}

func c68() {
	filicity := hr.Person{Name: "filicity", Age: 23}
	oliver := hr.Person{Name: "Oliver", Age: 25}

	people := []hr.Person{filicity, oliver}

	const greetPerson = `Hi I am {{.Name | upperCase }}. {{.Age}} years old {{"\n"}}`
	const greetPeople = `{{range .}}Hi I am {{.Name}}. {{.Age}} years old{{"\n"}}{{end}}`

	maps := make(template.FuncMap)
	maps["upperCase"] = upperCase

	gt := template.New("greetingFromPerson")
	gt.Funcs(maps)

	greetTemplate := template.Must(gt.Parse(greetPerson))
	gpt := template.New("greetPeopleTemplate")
	greetPeopleTemplate := template.Must(gpt.Parse(greetPeople))

	greetTemplate.Execute(os.Stdout, filicity)
	greetTemplate.Execute(os.Stdout, oliver)
	greetPeopleTemplate.Execute(os.Stdout, people)

}

func c69() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}
	todoDecoder := json.NewDecoder(resp.Body)
	todos := Todo{}
	todoDecoder.Decode(&todos)
	resp.Body.Close()

	indexTemplateByte, err := ioutil.ReadFile("D:/Project/@Site_Project/My-go-demy/src/sandbox3/index.html")
	if err != nil {
		return
	}
	indexTemplate := template.Must(template.New("index").Parse(string(indexTemplateByte)))
	indexTemplate.Execute(os.Stdout, todos)

}
