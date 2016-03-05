package main

import "fmt"

func main() {
	fmt.Print(getCode())
}

func getCode() string {
	code :=`package main

import "fmt"

func main() {
	fmt.Print(getCode())
}

func getCode() string {
	code :=
	return code[:99] + "\u0060" + code + "\u0060" + code[99:]
}`
	return code[:99] + "\u0060" + code + "\u0060" + code[99:]
}