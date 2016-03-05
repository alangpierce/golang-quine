package main

import "fmt"

func main() {
	code :=`package main

import "fmt"

func main() {
	code :=
	fmt.Print(code[:50] + "\u0060" + code + "\u0060" + code[50:])
}`
	fmt.Print(code[:50] + "\u0060" + code + "\u0060" + code[50:])
}