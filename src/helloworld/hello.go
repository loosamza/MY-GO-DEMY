package main

import (
	"fmt"
	"log"
	"myFmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(stringutil.Reverse("Hello World"))
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	myFmt.Println("Hello World")
	log.Println("Hello World")

}
