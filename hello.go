package main

import (
	"fmt"
	"time"
)

func fizzBuzz() {
	for i := 0; i < 256; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}
		fmt.Printf("%b, %x,%v\n", i, i, i)
	}
}
func scanString(question string) string {
	var name string
	fmt.Print(question)
	fmt.Scanf("%s", &name)
	return name
}

type function func(any, any) any

type object struct {
	numberProp int
	stringProp string
	boolean    bool
	fn         function
}

func main() {
	res := scanString("What is your name? ")
	fmt.Println("Hello, ", res)
	timer := time.Now()
	myslice := []int{}
	fizzBuzz()
	myslice = append(myslice, 3, 5, 6)
	for val, i := range myslice {
		fmt.Println(val, i)
	}
	// file, err := os.ReadFile("hello.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	myObject := object{
		numberProp: 17,
		stringProp: "Ayomide",
		boolean:    true,
		fn: func(a1, a2 any) any {
			fmt.Println("Hey I am", a1)
			return a2
		},
	}
	fmt.Printf("%v %v %v ", myslice, timer, "\n")
	fmt.Println(myObject.stringProp)
}
