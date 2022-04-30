package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"time"
)

// section 7 Function
func main() {

	fmt.Println("-----Introduction to function -----")
	c70()
	fmt.Println("-----recursive function -----")
	c71()
	fmt.Println("-----function and error relationship-----")
	c74()
	fmt.Println("-----function value-----")
	c75()
	fmt.Println("-----anonymous function concept-----")
	c76()
	fmt.Println("-----anonymous function concept challenge-----")
	c81_challenge()
	fmt.Println("---high order function and anonymous-----")
	c83()
	fmt.Println("---variadic function-----")
	c84()
	fmt.Println("---deferred function call-----")
	c85()

}

func avg(a, b float64) float64 {
	return (a + b) / 2
}
func avg2(a, b float64, c int64, d, f int) (float64, string, int) {
	return a, strconv.FormatInt(c, 10), d
}

func c70() {

	a := avg(1, 2)
	b := avg(3, 4)
	c, d, f := avg2(99.0, 0.01, 999, 5, 10)

	fmt.Println(a, b, c, d, f)
	fmt.Println(reflect.TypeOf(avg2))

}

type Node struct {
	value string
	nodes []Node
}

func c71() {
	p := Node{value: "p"}
	g := Node{value: "g"}
	b := Node{value: "b", nodes: []Node{p, g}}
	q := Node{value: "q"}
	s := Node{value: "s"}
	k := Node{value: "k"}
	r := Node{value: "r", nodes: []Node{q}}
	a := Node{value: "a", nodes: []Node{r, s}}
	root := Node{value: "a", nodes: []Node{b, a, k}}
	result := outline([]string{}, &root)
	fmt.Println(result)

}

func outline(stack []string, n *Node) [][]string {
	stack = append(stack, n.value)
	result := [][]string{}
	if len(n.nodes) == 0 {
		result = append(result, stack)
	}
	for _, v := range n.nodes {
		result = append(result, outline(stack, &v)...)
	}

	return result

}

const dbReady = false
const balance = 200
const numberToSuccess = 1

func connectDB(nTry int) error {
	if nTry == numberToSuccess {
		return nil
	}
	return errors.New("busy")
}

func getBalance() (int, error) {
	if !dbReady {
		if err := waitForDatabase(); err != nil {
			return 0, fmt.Errorf("getbalance: %v", err)
		}
	}
	return balance, nil
}

func waitForDatabase() error {
	timeout := 3 * time.Second
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		err := connectDB(tries)
		if err == nil {
			return nil
		}
		log.Printf("database is not responding (%v); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("waitForDatabase: timeout %v", timeout)
}

func withdraw(w int) (int, error) {

	balance, err := getBalance()
	if err != nil {
		return 0, fmt.Errorf("withdraw : %v", err)
	}
	if w > balance {
		return 0, errors.New("withdraw : insufficent fund")
	}
	return w, nil
}
func c74() {
	amount, err := withdraw(200)
	if err != nil {
		fmt.Println("main : ", err)
		os.Exit(1)
		// return
	}
	fmt.Println("Please collect your money : ", amount)
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func apply(a, b float64, op func(float64, float64) float64) (float64, error) {
	if op == nil {
		return math.NaN(), fmt.Errorf("apply : nil operation")
	}
	return op(a, b), nil

}

func c75() {
	a, _ := apply(1, 2, add)
	b, _ := apply(1, 2, sub)
	c, _ := apply(1, 2, nil)

	fmt.Println(a, b, c)

}

func createF() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}

func c76() {
	f := createF()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func createG(list []int) []func() {
	result := []func(){}
	for _, v := range list {
		var x = v
		result = append(result, func() {
			fmt.Println(x)
		})
	}
	return result
}
func c81_challenge() {
	gList := createG([]int{108, 12, 30, 4, 5})
	for _, v := range gList {
		v()
	}

}

func createFib() func(int) []int {
	fList := []int{0, 1, 1, 2, 3, 5}
	return func(n int) []int {
		if n > len(fList) {
			for n > len(fList) {
				lastIndex := len(fList) - 1
				fList = append(fList, fList[lastIndex]+fList[lastIndex-1])
			}
		}
		return fList
	}
}

func profileTime(f func(int) []int) func(int) []int {
	return func(a int) []int {
		start := time.Now()
		result := f(a)
		fmt.Printf("call with %d tooks %d ns\n", a, time.Now().Sub(start))
		return result
	}
}

func c83() {
	fib := createFib()
	fib = profileTime(fib)
	fib(60)
	fib(600000)
	fib(6000000)
}

func printEachLine(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func c84() {
	printEachLine(1, 2, "sss", "eee", '1')
	x := []interface{}{1, 2, "sss", "eee", '1'}
	printEachLine(x...)
}

func stopWatch(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("function %s : took time %v\n", name, time.Now().Sub(start).Seconds())
	}
}

func loadingImage() {
	defer stopWatch("loadingImage")()
	time.Sleep(3 * time.Second)
	fmt.Println("loadingImage : done")
}

func c85() {
	// defer
	x := 1
	addX := func(a int) int {
		x = x + a
		return x
	}
	defer fmt.Println(addX(3))
	fmt.Println(x)

	loadingImage()
}
