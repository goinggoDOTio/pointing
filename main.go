package main

import "fmt"

func main() {
	a := 50
	b := "Some string"
	c := 33

	m := &a
	n := &b
	o := &c

	fmt.Println("a value is", a)
	fmt.Println("m value is", *m)
	a = 66
	fmt.Println("Now m value is", *m)
	fmt.Println("a address is", &a)
	fmt.Println("n value is", *n)
	fmt.Println("b address is", &b)
	fmt.Println("o value is", *o)
	o = &a
	fmt.Println("Now o value is same as a:", *o)
}
